import { ActionContext, Commit } from "vuex";

export interface UserData {
	id: number,
	username: string,
	fullname: string,
}
export interface UserState {
	loading: boolean,
	error: Error | null,
	data: UserData | null,
}

const state: UserState = {
  loading: false,
  error: null,
  data: null,
};
const mutationTypes = {
  LOGIN_MAIN_REQUEST: "LOGIN_MAIN_REQUEST",
  LOGIN_MAIN_SUCCESS: "LOGIN_MAIN_SUCCESS",
  LOGIN_MAIN_FAILURE: "LOGIN_MAIN_FAILURE",

  GET_USER_REQUEST: "GET_USER_REQUEST",
  GET_USER_SUCCESS: "GET_USER_SUCCESS",
  GET_USER_FAILURE: "GET_USER_FAILURE",

	GET_WS_TOKEN_REQUEST: "GET_WS_TOKEN_REQUEST",
  GET_WS_TOKEN_SUCCESS: "GET_WS_TOKEN_SUCCESS",
  GET_WS_TOKEN_FAILURE: "GET_WS_TOKEN_FAILURE",
};
const mutations = {
	LOGIN_MAIN_REQUEST(state: UserState) {
		state.loading = true
		state.error = null
		state.data = null
	},
	LOGIN_MAIN_SUCCESS(state: UserState) {
		state.loading = false
	},
	LOGIN_MAIN_FAILURE(state: UserState, error: Error) {
		state.loading = false
		state.error = error
	},
	GET_USER_REQUEST(state: UserState) {
		state.loading = true
		state.error = null
		state.data = null
	},
  GET_USER_SUCCESS(state: UserState, data: UserData) {
		state.loading = false
		state.data = data
	},
  GET_USER_FAILURE(state: UserState, error: Error) {
		state.loading = false
		state.error = error
	},
	GET_WS_TOKEN_REQUEST(state: UserState){
		state.loading = true
		state.error = null
	},
  GET_WS_TOKEN_SUCCESS(state: UserState){
		state.loading = false
	},
  GET_WS_TOKEN_FAILURE(state: UserState, error: Error){
		state.loading = false
		state.error = error
	},
}
const actions = {
	async loginMainApp({ state, commit } : { state: UserState, commit: Commit}, loginData: any) {
		try {
			commit(mutationTypes.LOGIN_MAIN_REQUEST)
			const res = await window.axiosAuth.post('/authentication/main', loginData)
			if (res.status === 201) {
				localStorage.setItem('main_token', (res.data as any).main_token)
				commit(mutationTypes.LOGIN_MAIN_SUCCESS)
			}
		} catch(error) {
			commit(mutationTypes.LOGIN_MAIN_FAILURE, error)
		}
	},
	async getCurrentUser({ state, commit } : { state: UserState, commit: Commit}) {
		try {
			commit(mutationTypes.GET_USER_REQUEST)
			const res = await window.axios.get('/user')
			if (res.status === 200) {
				commit(mutationTypes.GET_USER_SUCCESS, res.data)
			}
		} catch(error) {
			commit(mutationTypes.GET_USER_FAILURE, error)
		}
	},
	async getWebsocketToken({ state, commit } : { state: UserState, commit: Commit}) {
		try {
			commit(mutationTypes.GET_WS_TOKEN_REQUEST)
			const res = await window.axiosAuth.post('/token/ws')
			if (res.status === 201) {
				localStorage.setItem('ws_token', (res.data as any).ws_token)
				commit(mutationTypes.GET_WS_TOKEN_SUCCESS)
			}
		} catch(error) {
			commit(mutationTypes.GET_WS_TOKEN_FAILURE, error)
		}
	}
};

export const userState = {
	state,
	mutations,
	actions,
}