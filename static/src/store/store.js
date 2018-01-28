import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ws: null
    },
    mutations: {
        startWS(state) {
            state.ws = new WebSocket("ws://" + location.host + "/ws");
            state.ws.onopen = event => {
                console.log(event);
            };
        }
    }
})