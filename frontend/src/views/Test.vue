<template>
  <div class="test">
    <h1>{{topic}}</h1>
    <div v-if="questionsLoaded">
      <Question :content="question" ref="question" @failed="failed" @solved="solved"/>
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

      hist:[-1],
      hakuna:[1],

      diffscale:0.2,//how much does the difference in difficulty get accounted for
      reward:1,
      penalty:0.8,

    }
  },
  mounted(){
    console.log('hakuna: ',JSON.stringify(this.hakuna));
    if (localStorage[this.topic]==null){
      console.log("no data on topic ",this.topic);
      localStorage[this.topic]=0
    }else{
      console.log(localStorage[this.topic])
    }
    console.log("skill level:",localStorage[this.topic]);
    var rating=Math.round(localStorage[this.topic])
    console.log('hist: ',this.hist);
    this.getNewQuestion(this.topic,rating)
  },
  methods:{

    getTopicInfo(topic){
      var data={"op":"GETINFO","topic":topic,"no":0}
      axios({method:"POST",url:"/api/getTopicInfo",data:data,headers:{"content-type":"text/plain"}}).then(result=>{
        localStorage[topic+"Info"]=JSON.stringify(result.data);
      })
    },

    getNewQuestion(topic,rating){
      var data={'topic':topic,'rating':rating,'history':this.hist}
      console.log('sending data:',data);
      axios({method:'POST',url:'api/new',data:data,headers:{'content-type':'text/plain'}}).then(result =>{
        this.question=result.data
        this.questionsLoaded=true
        console.log('receved question:',this.question);

      }).catch(error=>{
        console.log(error);
      })
    },
    solved(){
      console.log("difficulty:",this.question.rating);
      this.rating=parseFloat(localStorage[this.topic])
      console.log("start rating:",this.rating);
      var difference=this.question.rating-this.rating
      var change = Math.max(0,this.reward+difference*this.diffscale)
      this.rating+=change
      localStorage[this.topic]=this.rating
      console.log("new rating",this.rating);
      this.$refs.next.focus()


    },
    failed(){
      console.log('failed');
      console.log('difficulty:',this.question.rating);
      this.rating=parseFloat(localStorage[this.topic])
      console.log('start rating:',this.rating);
      var difference=this.question.rating=this.rating
      var change = Math.max(0,this.penalty+difference*this.diffscale)
      this.rating-=change
      localStorage[this.topic]=this.rating
      console.log('new rating:',this.rating);
      this.$refs.next.focus()

    },
    resetSkill(){
      localStorage[this.topic]=0
    },


    next(){
      console.log("question:", this.question);
      console.log("id:",this.question.id);
      this.hist.push(this.question.id)
      if (this.hist.length>10){
        this.hist=this.hist.slice(-10)
      }
      console.log("history:",this.hist);

      console.log("skill level:",localStorage[this.topic]);
      var rating=Math.round(localStorage[this.topic])
      // var hist = JSON.stringify(this.hist)
      this.getNewQuestion(this.topic,rating)
      this.$refs.question.start()
    },

  }
}
</script>
