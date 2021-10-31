import { createStore } from 'vuex'
import { clanState } from './clan'
import { userState } from './user'

export default createStore({
	modules: {
		user: {
			namespaced: true,
			...userState,
		},
		clan: {
			namespaced: true,
			...clanState,
		}
	},
})