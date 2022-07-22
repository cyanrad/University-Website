<script lang="ts">
import login from '@/components/LoginPanel.vue'

export default {
    components: {
        login
    }
}
</script>

<template>

    <body>
        <login />
    </body>
</template>

<style>
body {
    background-image: url(../../public/AU_Campus.jpg);
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
    background-position: center;
    font: 15px/16px "segoe ui", arial, sans-serif;
}
</style>