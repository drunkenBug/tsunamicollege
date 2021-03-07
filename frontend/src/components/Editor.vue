<template>
<div>
  <h1>Enter new Question</h1>
  <div class="whiteblock">
    <p>Title:
      <input type="text" v-model='title'>
    </p>
    <p>Solution:
      <input type="text" v-model='solution'>
    </p>
    <button type="submit" name="button" @click='submit'>
      Submit
    </button>

  </div>
</div>
</template>

<script>
import axios from 'axios'

export default{
    name:'Editor',
    components:{
    },
    data(){
      return {
        title:"",
        solution:"",
      }
    },
    methods:{
      submit(){
        var data = {"op":"SET","title":this.title,"solution":this.solution}
        axios({ method: "POST", url: "/api/setQuestion", data: data, headers: {"content-type": "text/plain" } }).then(result => {
            console.log("added new question")
            console.log(result.data)
        }).catch(error =>{
              console.error(error);
        })
      },

    }
  }
</script>

<style media="screen">
  button{
    background: #00f;
    border: 0;
    border-radius: 20px;
    color: white;
    margin-top: 20px;
    padding: 10px 20px;
  }
</style>
