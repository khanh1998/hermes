import { Store } from 'vuex'
import { ClanState } from '../store/clan';
import { MessageState } from '../store/message';
import { UserState } from '../store/user';

declare module '@vue/runtime-core' {
  // declare your own store states
  interface State {
    clan: ClanState,
    message: MessageState,
    user: UserState,
  }

  // provide typings for `this.$store`
  interface ComponentCustomProperties {
    $store: Store<State>
  }
}