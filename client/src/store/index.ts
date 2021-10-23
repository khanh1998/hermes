import { createStore } from 'vuex'
import { userState } from './user'

export default createStore({
	modules: {
		user: {
			namespaced: true,
			...userState,
		},
	},
})