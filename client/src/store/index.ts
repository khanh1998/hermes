import { createStore } from 'vuex'
import { clanState } from './clan'
import { userState } from './user'
import { messageState } from './message'

export default createStore({
	modules: {
		user: {
			namespaced: true,
			...userState,
		},
		clan: {
			namespaced: true,
			...clanState,
		},
		message: {
			namespaced: true,
			...messageState,
		}
	},
})