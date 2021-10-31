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
import { createNamespacedHelpers } from 'vuex';

interface LoginData {
	username: string
	password: string
}

const { mapActions } = createNamespacedHelpers('user')

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
		...mapActions([
			'loginMainApp',
			'getCurrentUser',
		]),
		async doLogin() {
			//
			console.log('login', this.loginData);
			await this.loginMainApp(this.loginData);
			// await this.getCurrentUser();
			await this.$router.push('/home');
		}
	}
})
</script>