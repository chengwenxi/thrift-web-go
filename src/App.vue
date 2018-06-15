<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
    //js
    import {Thrift} from "../gen-js/thrift"
    import {HelloThriftClient} from "../gen-js/HelloThrift"
    let transport = new Thrift.TXHRTransport("http://localhost:9092")
    let protocol = new Thrift.TJSONProtocol(transport)
    let client = new HelloThriftClient(protocol)
    client.SayHello("1234", function (result) {
        alert("result from go thrift server, " + result)
    })
    export default {
      name: 'app'
    }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
