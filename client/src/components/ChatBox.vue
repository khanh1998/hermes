<template>
  <div class="" v-if="!loading">
    <div
      class="flex flex-col border-green-500 border rounded h-full overflow-auto"
    >
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
import { UserData, UserState } from "../store/user";
import { Message, MessageState } from "../store/message";

interface ConvertedMessage {
  senderId: number;
  clanId: number;
  channelId: number;
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
    userState(): UserState {
      return this.$store.state.user;
    },
    userData(): UserData | null {
      return this.userState.data;
    },
    messageState(): MessageState {
      return this.$store.state.message;
    },
    messageData(): Array<Message> {
      return this.messageState.data;
    },
    loading(): boolean {
      return this.messageState.loading || this.userState.loading;
    },
    convertedMessages(): Array<ConvertedMessage> {
      const filterByChanelAndClan = this.messages.filter((mess: Message) => {
        return mess.channelId === this.channelId && mess.clanId === this.clanId;
      });
      return filterByChanelAndClan.map((mess: Message) => {
        const date = new Date(mess.time);
        const timeStr = `${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
        const side: Side =
          mess.senderId === this?.userData?.id ? Side.RIGHT : Side.LEFT;
        const showAvatar = this?.userData?.id !== mess.senderId;
        const letter = mess.senderId.toString().substr(0, 1);
        const isSender = !showAvatar;
        return { ...mess, timeStr, side, showAvatar, letter, isSender };
      });
    },
  },
  methods: {
    ...mapActions("user", ["getWebsocketToken", "getCurrentUser"]),
    getLastestMessages() {
      this.$store
        .dispatch("message/getLastestMessages", {
          clanId: this.clanId,
          channelId: this.channelId || 1,
        })
        .then(() => {
          this.messages = this.messageData;
        });
    },
    sendMessage() {
      if (this.message) {
        const message = this.makeMessage(this.message);
        this.ws.send(JSON.stringify(message));
        this.message = "";
      }
    },
    makeMessage(message: string): any {
      return {
        senderId: this?.userData?.id,
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
    },
  },
  watch: {
    channelId(newVal: any) {
      if (newVal && this.clanId) {
        this.getLastestMessages();
      }
    },
    clanId(newVal: any) {
      if (newVal && this.channelId) {
        this.getLastestMessages();
      }
    },
  },
  async created() {
    await this.getWebsocketToken();
    const token = localStorage.getItem("ws_token");
    console.log({ token });
    const host = import.meta.env.VITE_SOCKET_HOST;
    const ws = new WebSocket(`${host}?token=${token}`);
    this.ws = ws;
    ws.onopen = (ev: Event) => {
      const message = this.makeMessage(
        this.userData?.fullname + " | " + Date.now()
      );
      console.log({ open: ev });
      // ws.send(JSON.stringify(message));
      // setInterval(() => ws.send(JSON.stringify(message)), 1000);
    };
    ws.onmessage = (ev: MessageEvent<any>) => {
      console.log({ message: ev });
      const incomming: Message = JSON.parse(ev.data);
      if (this.messages.length > 10) {
        this.messages.shift();
        this.messages = [...this.messages];
      }
      this.messages.push(incomming);
    };
  },
});
</script>
