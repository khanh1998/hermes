<template>
	<div class="">
		<div class="flex flex-col border-green-500 border rounded">
			<div v-for="mess in convertedMessages" :key="mess.time" class="m-1 flex flex-row">
				<div class="rounded-full bg-green-900 w-8 h-8 text-center align-text mr-1">
					{{ mess.username.substr(0, 1) }}
				</div>
				<div class="flex flex-col items-end">
					<span class="border-green-600 bg-green-100 p-1 rounded">
						{{ mess.message }}
					</span>
					<span class="text-xs">{{ mess.timeStr }}</span>
				</div>
			</div>
		</div>
		<div class="grid grid-cols-12 mt-2 gap-x-1">
			<input
				v-model="message"
				type="text"
				class="focus:outline-none focus:ring focus:border-green-300 border border-green-600 col-span-10 rounded p-1"
				placeholder="your message"
			>
			<button
				@click="sendMessage"
				class="focus:outline-none focus:ring focus:border-green-300 col-span-2 border border-green-600 rounded bg-green-50 hover:bg-green-200"
			>
					Send
			</button>
		</div>
	</div>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import { mapActions, mapState } from "vuex";
interface Message {
	username: string,
	fullname: string,
	message: string,
	time: number,
}
interface ConvertedMessage {
	username: string,
	fullname: string,
	message: string,
	time: number,
	timeStr: string,
}
export default defineComponent({
	name: 'ChatBox',
	data: function () {
    return {
      ws: {} as WebSocket,
      messages: [] as Array<Message>,
			message: '',
    };
  },
	computed: {
		...mapState('user', [
			'data'
		]),
		convertedMessages(): Array<ConvertedMessage> {
			return this.messages.map((mess: Message) => {
				const date = new Date(mess.time)
				const timeStr = `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
				return { ...mess, timeStr };
			})
		},
	},
	methods: {
		...mapActions('user', [
			'getWebsocketToken',
			'getCurrentUser',
		]),
		sendMessage() {
			if (this.message) {
				this.ws.send(JSON.stringify(this.makeMessage(this.message)))
				this.message = ''
			}
		},
		makeMessage(message: string): any {
			return {
				name: this.data.fullname,
				username: this.data.username,
				message: message,
				time: Date.now(),
			}
		}
	},
  async created() {
		await this.getCurrentUser();
		await this.getWebsocketToken();
    const token = localStorage.getItem('ws_token');
    const host = import.meta.env.VITE_SOCKET_HOST;
    const ws = new WebSocket(`${host}?token=${token}`)
    this.ws = ws;
    ws.onopen = (ev: Event) => {
			const message = this.makeMessage('ping ping ping')
      console.log({ open: ev })
      ws.send(JSON.stringify(message))
			this.messages.push(message)
    }
    ws.onmessage = (ev: MessageEvent<any>) => {
      console.log({ message: ev})
			this.messages.push(JSON.parse(ev.data));
    }
  },
})
</script>