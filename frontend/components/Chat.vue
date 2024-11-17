<template>
    <div class="chat-container">
        <div class="chat-box" ref="chatBoxRef">
            <!-- Loop through messages and display them -->
            <div class="ai-person-container" v-for="(msg, index) in messages" :key="index">
                <div :class="msg.sender === 'user' ? 'person' : 'ai'">{{ msg.text }}</div>
            </div>
        </div>
        <div class="chat-input-container">
            <input class="chat-input" type="text" placeholder="Type a message..." v-model="message"
                @keyup.enter="sendMessage" />
            <button class="chat-submit" @click="sendMessage">Send</button>
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue';
import axios from 'axios';

// Create refs for message, messages, and chat box
const message = ref('');
const messages = ref([]);
const chatBoxRef = ref(null);

// Auto scroll to bottom when messages change
watch(() => messages.value.length, () => {
    setTimeout(() => {
        if (chatBoxRef.value) {
            chatBoxRef.value.scrollTop = chatBoxRef.value.scrollHeight;
        }
    }, 100); // Small delay to ensure content is rendered
});

// Define the sendMessage function
const sendMessage = async () => {
    // Trim and check if message is empty
    if (message.value.trim() === '') return;

    // Create a new user message
    const userMessage = {
        text: message.value.trim(),
        sender: 'user'
    };

    // Store the message text before clearing input
    const messageText = message.value;

    // Clear the input field immediately for better UX
    message.value = '';

    // Add user message to the chat box
    messages.value.push(userMessage);

    // Send the message to the backend API
    try {
        const response = await axios.post('http://127.0.0.1:8000/ask', {
            message: messageText
        });

        // Create a new AI message
        const aiMessageText = response.data?.response || 'No response from AI';
        const aiMessage = {
            text: aiMessageText,
            sender: 'ai'
        };

        console.log(aiMessageText);

        // Add AI message to the chat box
        messages.value.push(aiMessage);
    } catch (error) {
        console.error('Error sending message:', error);
    }

};



</script>

<style scoped>
.chat-container {
    display: flex;
    flex-direction: column;
    height: 100%;
}

.chat-box {
    flex: 1;
    overflow-y: auto;
    padding: 1rem;
}

.chat-input-container {
    display: flex;
    padding: 1rem;
    gap: 0.5rem;
    border-top: 1px solid #e9ecef;
}

.chat-input {
    flex: 1;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.chat-submit {
    padding: 0.5rem 1rem;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

/* Message styles for when you add messages */
.ai-person-container {
    margin-bottom: 1rem;
}

.person {
    background: #007bff;
    color: white;
    padding: 0.5rem 1rem;
    border-radius: 1rem;
    margin-left: auto;
    max-width: 80%;
    width: fit-content;
}

.ai {
    background: #e9ecef;
    color: black;
    padding: 0.5rem 1rem;
    border-radius: 1rem;
    margin-right: auto;
    max-width: 80%;
    width: fit-content;
}

.person-date,
.ai-date {
    font-size: 0.8rem;
    color: #6c757d;
    display: block;
    margin-top: 0.25rem;
}

.person-date {
    text-align: right;
}
</style>