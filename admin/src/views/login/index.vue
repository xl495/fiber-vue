<template>
    <div class="login-wrap">
        <div class="login-form">
            <login-form />
        </div>
    </div>
</template>

<script setup lang="ts">
import loginForm from './form.vue';

</script>

<style lang="scss">
.login-wrap {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #CDDCDC;
    background-image: radial-gradient(at 50% 100%, rgba(255,255,255,0.50) 0%, rgba(0,0,0,0.50) 100%), linear-gradient(to bottom, rgba(255,255,255,0.25) 0%, rgba(0,0,0,0.25) 100%);
    background-blend-mode: screen, overlay;

    .login-form {
        height: 300px;
        width: 400px;
        border-radius: 4px;
        border: 1px solid #4481eb;
        padding: 24px;
    }
}
</style>