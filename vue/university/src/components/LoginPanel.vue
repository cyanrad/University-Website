<script lang="ts">
import axios, { AxiosError } from "axios";
import { defineComponent } from "vue";

type loginResponse = {
    full_name: string;
}

export default defineComponent({
    data() {
        return {
            id: '',
            password: '',
            status: 0,
            fullName: ''
        }
    },

    methods: {
        async login(): Promise<void> {
            try {
                const { data } = await axios.post<loginResponse>(
                    'http://127.0.0.1:8080/users/login',
                    { id: this.id, password: this.password }
                )
                this.status = 2
                this.fullName = data.full_name
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    const err = error as AxiosError
                    if (err.response["status"] != 0) { this.status = 3 }
                    else {
                        this.status = 1
                    }
                }
                else {
                    console.log('unexpected error: ', error);
                }
            }
        },
    }

})
</script>

<template>
    <div class="form-box">
        <img src="/logo.png" alt="logo" class="logo">
        <h1>Welcome To Ajman University</h1>
        <hr class="line">

        <form>
            <p class="sub-text">Please enter your details to login</p>
            <p class="sub-text" style="color:red" v-if="status == 3">Incorrect Login</p>
            <p class="sub-text" style="color:green" v-if="status == 2">Hello {{ fullName }}</p>
            <p class="sub-text" style="color:black" v-if="status == 1">Servers Down</p>

            <input type="text" placeholder="Student ID" required v-model="id" />
            <input type="password" placeholder="Password" required v-model="password">
            <input type="button" class="login" @click="login" value="Login">
        </form>
    </div>
</template>

<style>
.form-box {
    width: 500px;
    background-color: #fffefe;
    margin: 12% auto;
    padding: 50px;
    box-shadow: 0 0 20px 2px rgba(0, 0, 0, 0.5)
}

h1 {
    text-align: center;
    margin-bottom: 40px;
    color: #336981;
}

.logo {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
    width: 150px;
    height: 150px
}

.line {
    border-top: 3px solid #336981;
    margin-top: -25px;
}

.sub-text {
    color: rgb(109, 109, 109);
    font-size: 20px;
    text-align: center;
}

form {
    display: flex;
    flex-direction: column;
    padding-top: 2rem;
    width: 100%;
}

form input {
    outline: none;
    padding: 0.8rem 1rem;
    margin-bottom: 0.8rem;
    font-size: 1.1rem;
}

form input:focus {
    border: 1.8px solid #336981;
}

form .login {
    outline: none;
    border: none;
    background: #336981;
    padding: 0.8rem 1rem;
    border-radius: 0.4rem;
    font-size: 1.1rem;
    color: #fff;
}

form .login:hover {
    background: #426f83;
    cursor: pointer;
}
</style>