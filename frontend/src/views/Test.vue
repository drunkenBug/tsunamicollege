<template>
  <div class="test">
    <h1>{{topic}}</h1>
    <div v-if="questionsLoaded">
      <Question :content="question" ref="question" @solved="solved"/>
    </div>
    <div class="next">
      <button class="test" type="button" ref="next" @click ="next">next</button>
    </div>
  </div>


</template>

<script>
import Question from '../components/Question.vue'
import axios from 'axios'
export default {
  name: 'Test',
  components: {
    Question
  },
  data(){
    return {
      questionsLoaded:false,
      question:{},
      data:[],
      index:1,
      topic:this.$route.params.id,

    }
  },
  mounted(){
    this.getQuestionById(this.topic,1)
    console.log(this.skillMap);
  },
  methods:{
    getQuestionById(topic,no){
      var data = {"op":"GET","topic":topic,"no":no}

      console.log(data)
      axios({ method: "POST", url: "/api/getQuestion", data: data, headers: {"content-type": "text/plain" } }).then(result => {
          console.log("receved question:")
          console.log(result.data)
          this.question=result.data
          this.questionsLoaded=true
          if (this.question.id==0){
            this.question.title="Congratulations, you completed the entrance test!"
          }
      }).catch(error =>{
            console.error(error);
      })
    },
    solved(){
      this.$refs.next.focus()
    },
    next(){
      this.index=this.index+1,
      this.getQuestionById(this.topic,this.index)
      this.$refs.question.start()
    },

  }
}
</script>
