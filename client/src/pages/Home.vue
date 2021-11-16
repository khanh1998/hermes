<template>
  <div class="grid grid-cols-12 gap-x-2 p-2 h-auto" v-show="!isLoading">
    <div class="col-span-3">
      <user-info :user="userData.data" class="mb-2" />
      <clan-list
        :clans="userClans"
        :selected-clan-id="selectedClanId"
        @change-clan="handleChangeClan"
        class="mb-2"
      />
      <channel-list
        :channels="clanChannels"
        :current-channel-id="selectedChannelId"
        @change-channel="handleChangeChannel"
      />
    </div>
    <div class="col-span-9 flex flex-col">
      <search-box class="mb-2" @search-for="doSearch" />
      <router-view></router-view>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { mapActions } from "vuex";
import ChannelList from "../components/ChannelList.vue";
import ChatBox from "../components/ChatBox.vue";
import { Channel, ClanState } from "../store/clan";
import { UserState, ClanShort } from "../store/user";
import { HOME_PAGE } from "../constants/constant";
import ClanList from "../components/ClanList.vue";
import UserInfo from "../atoms/UserInfo.vue";
import SearchBox from "../atoms/SearchBox.vue";
export default defineComponent({
  components: { ChatBox, ChannelList, ClanList, UserInfo, SearchBox },
  name: "Home",
  data() {
    return {
      channelId: HOME_PAGE.DEFAULT_CHANNEL_ID,
      clanId: HOME_PAGE.DEFAULT_CLAN_ID,
    };
  },
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
    selectedClanId(): number | null {
      if (this.clanId !== HOME_PAGE.DEFAULT_CLAN_ID) {
        return this.clanId;
      }
      if (this.userClans?.length === 0) {
        return null;
      } else {
        return this.userClans[0].id;
      }
    },
    selectedChannelId(): number | null {
      if (this.channelId !== HOME_PAGE.DEFAULT_CHANNEL_ID) {
        return this.channelId;
      }
      if (this.clanChannels.length === 0) {
        return null;
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
    handleChangeChannel(channelId: number) {
      this.channelId = channelId;
    },
    handleChangeClan(clanId: number) {
      this.clanId = clanId;
    },
    doSearch(searchString: any) {
      this.$router.push(`/search/${searchString}`)
    },
  },
  async created() {
    await this.getCurrentUser();
    await this.getClanById(this.selectedClanId);
    // await this.getWebsocketToken();
  },
});
</script>
