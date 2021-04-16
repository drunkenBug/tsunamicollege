<template>
  <div class="element">
    <button class='topicButton'>
      <router-link :to="{name:'Test',params:{id:topicID}}" >
        <h3> {{topicID}}</h3>
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
