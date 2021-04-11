<template>
  <div class="element">
    <button class='test'>
      <router-link :to="{name:'Test',params:{id:topicID}}" >
        <h2>✭✭✭☆☆ {{topicID}}</h2>
      </router-link>
    </button>
  </div>



</template>

<script>
  export default{
    props:[
      'topicID',
    ],
    data(){
      return {
        input:""
      }
    },
    methods:{
      created(){
        console.log("created");
      },
      get_link(){
        var base="Test/"
        var tail="plusminus"
        var link=base+tail
        console.log(link);
        return "Test/"+this.title
      }
    }
  }
</script>
<style media="screen">
  h2{
    color:#fff;
  },
  div{
    background-color:red;
  },
  button.element{
    background-color:red;
  }

</style>
