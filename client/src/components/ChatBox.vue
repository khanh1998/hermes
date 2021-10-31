<template>
  <div class="">
    <div class="flex flex-col border-green-500 border rounded">
      <template
        v-for="mess in convertedMessages"
        :key="mess.time"
        class="m-1 flex flex-row"
      >
        <message-item
          :message="mess.message"
          :time-str="mess.timeStr"
          :username="mess.letter"
          :side="mess.side"
          :show-avatar="mess.showAvatar"
          :is-sender="mess.isSender"
        />
      </template>
    </div>
    <div class="grid grid-cols-12 mt-2 gap-x-1">
      <input
        v-model="message"
        type="text"
        class="
          focus:outline-none focus:ring focus:border-green-300
          border border-green-600
          col-span-10
          rounded
          p-1
        "
        placeholder="your message"
        @keyup="pressEnterToSend"
      />
      <button
        @click="sendMessage"
        class="
          focus:outline-none focus:ring focus:border-green-300
          col-span-2
          border border-green-600
          rounded
          bg-green-50
          hover:bg-green-200
        "
      >
        Send
      </button>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import { mapActions, mapState } from "vuex";
import MessageItem from "../atoms/MessageItem.vue";
import { Side } from "../constants/Chatbox";
interface Message {
  senderId: number;
  clanId: number;
  message: string;
  time: number;
}
interface ConvertedMessage {
  senderId: number;
  clanId: number;
  message: string;
  time: number;
  timeStr: string;
  showAvatar: boolean;
  letter: string;
  isSender: boolean;
}

export default defineComponent({
  components: { MessageItem },
  name: "ChatBox",
  data: function () {
    return {
      ws: {} as WebSocket,
      messages: [] as Array<Message>,
      message: "",
    };
  },
  props: {
    clanId: Number,
    channelId: Number,
  },
  computed: {
    ...mapState("user", {
      userData: (state: any) => state.data,
    }),
    convertedMessages(): Array<ConvertedMessage> {
      return this.messages.map((mess: Message) => {
        const date = new Date(mess.time);
        const timeStr = `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
        const side: Side =
          mess.senderId === this.userData.id ? Side.RIGHT : Side.LEFT;
        const showAvatar = this.userData.id !== mess.senderId;
        const letter = mess.senderId.toString().substr(0,1);
        const isSender = !showAvatar;
        return { ...mess, timeStr, side, showAvatar, letter, isSender };
      });
    },
  },
  methods: {
    ...mapActions("user", ["getWebsocketToken", "getCurrentUser"]),
    sendMessage() {
      if (this.message) {
        const message = this.makeMessage(this.message);
        this.ws.send(JSON.stringify(message));
        this.messages.push(message);
        this.message = "";
      }
    },
    makeMessage(message: string): any {
      return {
        senderId: this.userData.id,
        clanId: this.clanId,
        channelId: this.channelId,
        message: message,
        time: Date.now(),
      };
    },
    pressEnterToSend(ev: KeyboardEvent) {
      if (ev.key === "Enter") {
        this.sendMessage();
      }
    }
  },
  async created() {
    await this.getWebsocketToken();
    const token = localStorage.getItem("ws_token");
    console.log({ token });
    const host = import.meta.env.VITE_SOCKET_HOST;
    const ws = new WebSocket(`${host}?token=${token}`);
    this.ws = ws;
    ws.onopen = (ev: Event) => {
      const message = this.makeMessage("ping ping ping");
      console.log({ open: ev });
      ws.send(JSON.stringify(message));
      this.messages.push(message);
    };
    ws.onmessage = (ev: MessageEvent<any>) => {
      console.log({ message: ev });
      this.messages.push(JSON.parse(ev.data));
    };
  },
});
</script>
