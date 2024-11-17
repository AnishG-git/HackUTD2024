<template>
    <div v-html="markdownContent" class="markdown-body" style="overflow-y: auto; max-height: 100vh;"></div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import axios from 'axios';
import MarkdownIt from 'markdown-it';
import hljs from 'highlight.js'; // Optional: for code highlighting

// Initialize markdown-it with options
const md = new MarkdownIt({
    html: true,        // Enable HTML tags in source
    xhtmlOut: true,    // Use '/' to close single tags (<br />)
    breaks: true,      // Convert '\n' in paragraphs into <br>
    linkify: true,     // Autoconvert URL-like text to links
    typographer: true, // Enable some language-neutral replacement + quotes beautification
    highlight: function (str, lang) {
        if (lang && hljs.getLanguage(lang)) {
            try {
                return hljs.highlight(str, { language: lang }).value;
            } catch (__) {}
        }
        return ''; // use external default escaping
    }
});

// Props definition
const props = defineProps({
    filePath: {
        type: String,
        required: true
    }
});

// Reactive reference for markdown content
const markdownContent = ref('');

// Load markdown content
const loadMarkdown = async () => {
    try {
        const response = await axios.get(props.filePath);
        markdownContent.value = md.render(response.data);
    } catch (error) {
        console.error('Error loading markdown file:', error);
    }
};

// Watch for filePath changes
watch(() => props.filePath, () => {
    loadMarkdown();
});

// Load markdown on mount
onMounted(() => {
    loadMarkdown();
});
</script>

<style>
/* Basic styles for markdown content */
.markdown-body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif;
    font-size: 16px;
    line-height: 1.5;
    word-wrap: break-word;
    padding: 2rem;
}

/* Headers */
.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
    margin-top: 24px;
    margin-bottom: 16px;
    font-weight: 600;
    line-height: 1.25;
}

/* Links */
.markdown-body a {
    color: #0366d6;
    text-decoration: none;
}

.markdown-body a:hover {
    text-decoration: underline;
}

/* Code blocks */
.markdown-body pre {
    padding: 16px;
    overflow: auto;
    font-size: 85%;
    line-height: 1.45;
    background-color: #f6f8fa;
    border-radius: 6px;
}

.markdown-body code {
    padding: 0.2em 0.4em;
    margin: 0;
    font-size: 85%;
    background-color: rgba(27,31,35,0.05);
    border-radius: 6px;
}

/* Lists */
.markdown-body ul,
.markdown-body ol {
    padding-left: 2em;
}

/* Blockquotes */
.markdown-body blockquote {
    padding: 0 1em;
    color: #6a737d;
    border-left: 0.25em solid #dfe2e5;
    margin: 0;
}

/* Tables */
.markdown-body table {
    border-spacing: 0;
    border-collapse: collapse;
    margin-bottom: 16px;
}

.markdown-body table th,
.markdown-body table td {
    padding: 6px 13px;
    border: 1px solid #dfe2e5;
}

.markdown-body table tr {
    background-color: #fff;
    border-top: 1px solid #c6cbd1;
}

.markdown-body table tr:nth-child(2n) {
    background-color: #f6f8fa;
}

/* Images */
.markdown-body img {
    max-width: 100%;
    box-sizing: border-box;
}
</style>