import { Commit } from "vuex";
export interface Message {
  senderId: number;
  clanId: number;
  channelId: number;
  message: string;
  time: number;
}
export interface MessageState {
  loading: boolean;
  error: Error | null;
  data: Array<Message>;
}

const state: MessageState = {
  loading: false,
  error: null,
  data: [],
};
const mutationTypes = {
  GET_LATEST_MESSAGE_REQUEST: "GET_LATEST_MESSAGE_REQUEST",
  GET_LATEST_MESSAGE_SUCCESS: "GET_LATEST_MESSAGE_SUCCESS",
  GET_LATEST_MESSAGE_FAILURE: "GET_LATEST_MESSAGE_FAILURE",
};
const mutations = {
  GET_LATEST_MESSAGE_REQUEST(state: MessageState) {
    state.loading = true;
    state.error = null;
  },
  GET_LATEST_MESSAGE_SUCCESS(state: MessageState, messages: Array<Message>) {
    state.data = [...state.data, ...messages];
    state.loading = false;
  },
  GET_LATEST_MESSAGE_FAILURE(state: MessageState, error: Error) {
    state.error = error;
    state.loading = false;
  },
};
const actions = {
  async getLastestMessages(
    { state, commit }: { state: MessageState; commit: Commit },
    params: { clanId: number; channelId: number }
  ) {
    try {
      commit(mutationTypes.GET_LATEST_MESSAGE_REQUEST);
      const res = await window.axios.get(
        `/message/latest?clan=${params.clanId}&channel=${params.channelId}&limit=10`
      );
			let messages = (res as any).data.body.hits.hits;
			messages = messages.reverse().map((mess: any) => mess._source)
			console.log(messages);
      commit(mutationTypes.GET_LATEST_MESSAGE_SUCCESS, messages);
    } catch (error) {
      commit(mutationTypes.GET_LATEST_MESSAGE_FAILURE, error);
    }
  },
};

export const messageState = {
  state,
  mutations,
  actions,
};
