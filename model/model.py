import numpy as np
import tensorflow as tf
from tensorflow.keras.datasets import imdb
from tensorflow.keras.preprocessing.sequence import pad_sequences
from tensorflow.keras.preprocessing.text import Tokenizer
from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn

app = FastAPI()
model = tf.keras.models.load_model('model.keras')

# Prepare tokenizer
tokenizer = imdb.get_word_index()
word_index = {k: (v + 3) for k, v in tokenizer.items()}
word_index["<PAD>"] = 0
word_index["<START>"] = 1
word_index["<UNK>"] = 2
word_index["<UNUSED>"] = 3

def preprocess_review(review):
    review = review.lower()
    tokens = review.split()
    sequence = [word_index.get(word, 2) for word in tokens]
    return pad_sequences([sequence], maxlen=500, padding='post')

# Pydantic model for input validation
class Review(BaseModel):
    review: str

@app.post("/predict")
def predict_sentiment(review: Review):
    print(review)
    try:
        # Preprocess the input review
        padded_review = preprocess_review(review.review)

        # Make a prediction
        prediction = model.predict(np.array(padded_review))

        # Convert the prediction to a class
        predicted_class = "Positive" if prediction[0] > 0.5 else "Negative"

        return {"review": review.review, "prediction": predicted_class}
    
    except Exception as e:
        # return error if any exception occurs
        return {"error": "Internal server error", "details": str(e)}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)