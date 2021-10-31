import { Commit } from "vuex";
export interface Channel {
	id: number,
	name: string,
}
export interface ClanData {
  id: 1;
  domain: string;
  name: string;
  chiefId: number;
	chief: { id: number, name: string },
	channels: Array<Channel>
}

export interface ClanState {
	loading: boolean,
	error: Error | null,
	data: ClanData | null
}

const state: ClanState = {
  loading: false,
  error: null,
	data: null
};
const mutationTypes = {
	GET_CLAN_REQUEST: 'GET_CLAN_REQUEST',
	GET_CLAN_SUCCESS: 'GET_CLAN_SUCCESS',
	GET_CLAN_FAILURE: 'GET_CLAN_FAILURE'
};
const mutations = {
	GET_CLAN_REQUEST(state: ClanState) {
		state.loading = true;
		state.data = null;
		state.error = null;
	},
	GET_CLAN_SUCCESS(state: ClanState, data: ClanData) {
		state.data = data;
		state.error = null;
		state.loading = false;
	},
	GET_CLAN_FAILURE(state: ClanState, error: Error) {
		state.data = null;
		state.error = error;
		state.loading = false;
	}
};
const actions = {
	async getClanById({ state, commit } : { state: ClanState, commit: Commit}, clanId: number) {
		try {
			commit(mutationTypes.GET_CLAN_REQUEST)
			const res = await window.axios.get(`/clan/${clanId}`)
			if (res.status >= 200 && res.status < 300) {
				commit(mutationTypes.GET_CLAN_SUCCESS, res.data)
			} else {
				throw new Error('fail due to response status')
			}
		} catch (error) {
			commit(mutationTypes.GET_CLAN_FAILURE, error)	
		}
	}
};

export const clanState = {
  state,
  mutations,
  actions,
};
