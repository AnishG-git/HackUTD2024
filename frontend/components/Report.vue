<template>
    <div v-html="markdownContent" class=""></div>
</template>

<script>
import axios from 'axios';
import { marked } from 'marked';

export default {
    data() {
        return {
            markdownContent: ''
        };
    },
    props: {
        filePath: {
            type: String,
            required: true
        }
    },
    mounted() {
        this.loadMarkdown();
    },
    methods: {
        async loadMarkdown() {
            try {
                const response = await axios.get(this.filePath);
                this.markdownContent = marked(response.data);
            } catch (error) {
                console.error('Error loading markdown file:', error);
            }
        }
    }
};
</script>

<style scoped>
/* Add any styles you need here */
</style>