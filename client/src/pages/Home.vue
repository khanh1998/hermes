<template>
  <div class="grid grid-cols-12 gap-x-2 p-2" v-if="!isLoading">
    <channel-list
      class="col-span-3"
      :channels="clanChannels"
      :current-channel-id="selectedChannelId"
    />
    <chat-box
      class="col-span-9"
      :channel-id="selectedChannelId"
      :clan-id="selectedClanId"
    />
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { mapActions, mapState } from "vuex";
import ChannelList from "../components/ChannelList.vue";
import ChatBox from "../components/ChatBox.vue";
import { Channel, ClanData, ClanState } from "../store/clan";
import { UserData, UserState, ClanShort } from "../store/user";
export default defineComponent({
  components: { ChatBox, ChannelList },
  name: "Home",
  computed: {
    isLoading(): boolean {
      return !!this?.userData?.loading && !!this?.clanData?.loading;
    },
    userData(): UserState | null {
      return this.$store.state.user;
    },
    userClans(): Array<ClanShort> {
      return this.userData?.data?.clans || [];
    },
    selectedClanId(): number {
      if (this.userClans?.length === 0) {
        throw new Error("");
      } else {
        return this.userClans[0].id;
      }
    },
    selectedChannelId(): number {
      if (this.clanChannels.length === 0) {
        throw new Error("channel empty");
      } else {
        return this.clanChannels[0].id;
      }
    },
    clanData(): ClanState | null {
      return this.$store.state.clan;
    },
    clanChannels(): Channel[] {
      return this.clanData?.data?.channels || [];
    },
  },
  methods: {
    ...mapActions("user", ["getCurrentUser", "getWebsocketToken"]),
    ...mapActions("clan", ["getClanById"]),
  },
  async created() {
    await this.getCurrentUser();
    await this.getClanById(this.selectedClanId);
    await this.getWebsocketToken();
  },
});
</script>
