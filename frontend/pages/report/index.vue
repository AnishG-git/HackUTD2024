<template>
    <div class="flex flex-col w-screen h-screen bg-gray-900 text-white">
        <navbar />
        <div class="flex h-full">
            <div
                class="m-4 max-w-screen-lg flex text-black text-md bg-white rounded-lg shadow-md border-blue-600 border-4">
                <Report filePath="/business_report.md" />
            </div>
            <div class="mt-32 w-3/6 h-4/6 text-black text-md bg-white rounded-lg shadow-md border-blue-600 border-4">
                <Chat>

                    
                </Chat>
            </div>


        </div>
    </div>
</template>

<script>
import Report from '@/components/Report.vue';
export default {
    name: 'ReportPage',
    data() {
        return {
            // Add your data properties here
        };
    },
    methods: {
        // Add your methods here
    },
    mounted() {
        // Add your mounted logic here
    }
};
</script>