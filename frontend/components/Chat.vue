<template>
    <div class="chat-container">
        <div class="chat-box">
            <div class="ai-person-container">
                <div class="person">Hello! How can I help API's today?</div>
            </div>
        </div>
        <div class="chat-input-container">
            <input class="chat-input" type="text" placeholder="Type a message..." />
            <button class="chat-submit">Send</button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ChatComponent'
}
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