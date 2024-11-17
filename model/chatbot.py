from pydantic import BaseModel
from fastapi import FastAPI, HTTPException
import json
import openai
import pandas as pd
import uvicorn
import os

app = FastAPI()

def saveCSV(file_path):
    with open(file_path, 'r') as file:
        data = json.load(file)
    print(data)
    df = pd.DataFrame(data)
    
    csv_file_path = 'data.csv'
    df.to_csv(csv_file_path, index=False)
    print(f"CSV file created at: {csv_file_path}")
    return csv_file_path

def generate_api_insights(csv_file_path, output_file_path="insights.json"):
    """
    Generate API insights from a CSV file and save them to a JSON file.

    Parameters:
        csv_file_path (str): Path to the input CSV file containing API data.
        output_file_path (str): Path to save the output insights JSON file.
    """
    # Initialize SambaNova LLM client
    client = openai.OpenAI(api_key="9cc5e443-487d-450c-913d-2d027b5ea1eb", base_url="https://api.sambanova.ai/v1")

    # Load the CSV file
    data = pd.read_csv(csv_file_path)

    # Ensure only the relevant columns are used
    columns_to_analyze = [
        'sdk_key', 'request_id', 'raw_req_id', 'raw_resp_id', 
        'req_in', 'req_out', 'endpoint', 'resp_states', 'method'
    ]
    data = data[columns_to_analyze]
    data['req_in'] = pd.to_numeric(data['req_in'], errors='coerce')
    data['req_out'] = pd.to_numeric(data['req_out'], errors='coerce')
    # Function to get insights from SambaNova LLM
    def get_llm_insights(prompt):
        response = client.chat.completions.create(
            model="Meta-Llama-3.1-8B-Instruct",
            messages=[{"role": "system", "content": prompt}],
            temperature=0.7,
            top_p=0.9
        )
        return response.choices[0].message.content

    # Function to summarize metrics for an endpoint
    def summarize_metrics(group):
        total_requests = len(group)
        unique_methods = group['method'].value_counts().to_dict()
        response_states = group['resp_states'].value_counts().to_dict()
        avg_request_processing_time = (group['req_out'] - group['req_in']).mean()
    
        # Summarize data
        summary = (
            f"Total requests: {total_requests}.\n"
            f"Unique methods used: {unique_methods}.\n"
            f"Response states distribution: {response_states}.\n"
            f"Average request processing time: {avg_request_processing_time:.2f} ms."
        )
        return summary

    # Generate insights for each endpoint
    insights = []

    for endpoint, group in data.groupby("endpoint"):
        # Summarize the metrics
        summary = summarize_metrics(group)

        # Use SambaNova LLM to generate analysis and suggestions
        llm_response = get_llm_insights(
            f"You are an API analytics expert. Given the following metrics for the endpoint '{endpoint}', "
            f"provide an analysis and actionable suggestions:\n\n{summary}\n\n"
            "Respond in the format: 'Analysis: <analysis>. Suggestion: <suggestion>'."
        )

        # Split the response into analysis and suggestion
        analysis, suggestion = "", ""
        if "Analysis:" in llm_response and "Suggestion:" in llm_response:
            analysis = llm_response.split("Analysis:")[1].split("Suggestion:")[0].strip()
            suggestion = llm_response.split("Suggestion:")[1].strip()
        else:
            analysis = llm_response  # Fallback in case format is not strictly followed

        # Format the output
        insights.append({
            "endpoint": endpoint,
            "analysis": analysis,
            "suggestion": suggestion
        })

    # Save insights to a JSON file
    with open(output_file_path, "w") as f:
        json.dump(insights, f, indent=4)

    print(f"Insights saved to {output_file_path}")

def generate_business_report(df):
    """Create a prompt for generating a business report from the data"""
    # Calculate key metrics
    metrics = {
        "total_requests": len(df),
        "total_endpoints": df['endpoint'].nunique(),  # Number of unique endpoints
        "total_methods": df['method'].nunique(),  # Number of unique HTTP methods
        "successful_responses": len(df[df['resp_states'].astype(str).apply(lambda x: x.startswith('2'))]),  # Successful responses (assuming resp_states starts with '2' for success)
        "error_responses": len(df[df['resp_states'].astype(str).apply(lambda x: x.startswith(('4', '5')))]),  # Errors (responses starting with '4' or '5')
        "suspicious_requests": len(df[pd.to_numeric(df['req_in'], errors='coerce') > 10000]),  # Requests with suspiciously large 'req_in'
        "most_frequent_endpoint": df['endpoint'].mode().iloc[0],  # Most frequently accessed endpoint
        "date_range": f"{df['req_in'].min()} to {df['req_in'].max()}"  # Date range based on 'req_in' column (or another time-based column)
    } 


    
    report_prompt = f"""
    Based on the following API metrics data, generate a comprehensive business report with actionable insights:
    Make sure it's properly formatted in Markdown.

    Time Period: {metrics['date_range']}
    Total Requests: {metrics['total_requests']}
    Total Endpoints: {metrics['total_endpoints']}
    Total Methods: {metrics['total_methods']}
    Successful Responses: {metrics['successful_responses']}
    Error Responses: {metrics['error_responses']}
    Suspicious Requests (req_in > 10000): {metrics['suspicious_requests']}
    Most Frequent Endpoint: {metrics['most_frequent_endpoint']}
    
    Firstly, list the top 5 endpoints that need attention

    Please analyze this data and provide:
    1. Executive Summary
    2. Performance Analysis
    3. Security Concerns
    4. Recommendations
    Make sure to highlight any suspicious patterns or potential issues.
    """
    return report_prompt

class ChatRequest(BaseModel):
    message: str

@app.post("/report")
def chat():
    try:
        df = pd.read_csv("data.csv")
        prompt = generate_business_report(df)
        client = openai.OpenAI(api_key="9cc5e443-487d-450c-913d-2d027b5ea1eb", base_url="https://api.sambanova.ai/v1")
        response = client.chat.completions.create(
            model='Meta-Llama-3.1-8B-Instruct',
            messages=[
                {"role": "system", "content": "You are a business analyst specializing in API metrics analysis."},
                {"role": "user", "content": prompt}
            ],
            temperature=0.1,
            top_p=0.1
        )
        print(response)
        report = response.choices[0].message.content
        with open("business_report.md", "w") as f:
            f.write(report)
        print("Business report saved to business_report.md")
        return {"response": report}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e)) 

def process_question(df, question):
    """Prepare context for answering specific questions about the data"""
    context = f"""Based on the API metrics data, answer the following question: {question}

    Relevant metrics:
    - Total requests: {len(df)}
    - Success rate: {(df['resp_states'].astype(str).str.startswith('2').mean() * 100):.2f}%
    - Error count (4xx/5xx): {len(df[df['resp_states'].astype(str).str.startswith(('4', '5'))])}
    - Most common method: {df['method'].mode().iloc[0]}
    - Unique endpoints: {len(df['endpoint'].unique())}
    """
    return context

@app.post("/ask")
def chat(request: ChatRequest):
    try:
        df = pd.read_csv("data.csv")
        context = process_question(df, request.message)
        client = openai.OpenAI(api_key="9cc5e443-487d-450c-913d-2d027b5ea1eb", base_url="https://api.sambanova.ai/v1")
        response = client.chat.completions.create(
            model='Meta-Llama-3.1-8B-Instruct',
            messages=[
                {"role": "system", "content": "You are a business analyst specializing in API metrics analysis."},
                {"role": "user", "content": context}
            ],
            temperature=0.1,
            top_p=0.1
        )
        print(response)
        return {"response": response.choices[0].message.content}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/genData") 
def generateData():
    saveCSV("data.json")
    return {"message": "CSV file created successfully"}

@app.post("/genInsights")
def generateInsights():
    try:
        generate_api_insights("data.csv")
        return {"message": "Insights generated successfully"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    
if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)