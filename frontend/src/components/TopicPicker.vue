<template>
  <div class="bubble">

    <div class ='scrollbox'>
      <TopicButton topicID="Plus"/>
      <TopicButton topicID='Minus' />
      <TopicButton topicID="Ein mal eins"/>
      <TopicButton topicID="Geteilt"/>
      <TopicButton topicID="Brüche" />
      <TopicButton topicID='Prozente' />
      <TopicButton topicID="Funktionen"/>
      <TopicButton topicID="Koordinaten"/>
      <TopicButton topicID="Graphen"/>
    </div>
  </div>


</template>

<script>
import TopicButton from './TopicButton.vue'
  export default{
    components:{
      TopicButton,
    },
    props:[
      'title',
    ],
    data(){
      return {
        result:"...",
        input:""
      }
    },
    methods:{
      foo(){
      },
    }
  }
</script>
