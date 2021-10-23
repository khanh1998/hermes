<template>
	<div class="grid grid-cols-12">
		<div class="col-span-6 col-start-3 flex flex-col">
			<div class="flex flex-col">
				<span>User</span>
				<input type="text" placeholder="username" v-model="loginData.username"/>
			</div>
			<div class="flex flex-col">
				<span>Password</span>
				<input type="text" placeholder="password" v-model="loginData.password">
			</div>
			<button @click="doLogin">Login</button>
		</div>
	</div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'

interface LoginData {
	username: string
	password: string
}

export default defineComponent({
	data() {
		return {
			loginData: {
				username: '',
				password: '',
			} as LoginData
		}
	},
	methods: {
		async doLogin() {
			//
			console.log("login", this.loginData);
		}
	}
})
</script>