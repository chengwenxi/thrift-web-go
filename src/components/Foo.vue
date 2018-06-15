<template>
    <div>Foo</div>
</template>

<script>
    // //nodejs
    // import thrift from "thrift"
    // import helloThrift from "../../gen-nodejs/HelloThrift"
    // var transport1 = thrift.TBufferedTransport;
    // var protocol1 = thrift.TJSONProtocol;
    // var connection = thrift.createXHRConnection("localhost","9092",{
    //     transport : transport,
    //     protocol : protocol})
    // var client1 = thrift.createClient(helloThrift, connection);
    // client1.constructor.SayHello("1234",function (result) {
    //     console.log(result)
    // })

    export default {
        name: "Foo",
        created: function () {
            // // eslint-disable-next-line
            // var transport = new Thrift.TXHRTransport("http://localhost:9092")
            // // eslint-disable-next-line
            // var protocol = new Thrift.Protocol(transport)
            // // eslint-disable-next-line
            // var client = new HelloThriftClient(protocol)
            // client.send_Say("1234", function (result) {
            //     console.log(result)
            // })
        }
    }
</script>

<style scoped>


</style>