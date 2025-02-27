<script setup>

const props = defineProps({
    title: String,
    user: String,
    role: String,
});

</script>
<template>
    <div class="container">
        <div class="title">
            <img :src="`/assets/icon/${props.title}.png`" alt="">
            <h1>{{ props.title }}</h1>
        </div>
        <h1>Halo {{ props.role }}, {{ props.user }}</h1>
    </div>
</template>
<style scoped>
.container {
    margin-left: 0 !important;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.title {
    display: flex;
    align-items: center;
    gap: 1rem;
}

.container>h1 {
    font-weight: 500;
}
</style>