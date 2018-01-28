<template>
  <div id="app">
    <h1>OPEN DEV TOOLS TO SEE WEBSOCKET MESSAGES...</h1>
  </div>
</template>

<script>
import xs from "xstream";

export default {
  name: "app",
  data() {
    return {
      ws: null,

      producer: {
        start: listener => {
          this.$store.commit("startWS");
          this.$store.state.ws.onmessage = event => {
            console.log("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
            console.log(event);
            listener.next(event);
          };
        },
        stop: () => {
          this.$store.state.ws.onclose = event => {
            console.log("Stopped listening to websocket.");
          };
          this.$store.state.ws.close();
        }
      },
      listener: {
        next: value => {
          console.log(value);
        },
        error: err => {
          console.error("Error reading from websocket stream: ", err);
        },
        complete: () => {
          console.log("Stream complete.");
        }
      }
    };
  },
  computed: {
    main$() {
      return xs.createWithMemory(this.producer);
    },
    ping$() {
      return xs.from(this.main$);
    }
  },
  mounted() {
    this.ping$.addListener(this.listener);
  }
};
</script>

<style>

</style>