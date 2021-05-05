<template>
  <div class="bubble">
    <div class="question">

      <h2>{{content.title}}</h2>
      <input class="test" type="number" v-model='input' ref="input" @keyup="keyup" >
      <button class="test" type="button" name="button" @click='action' >
        <!-- <div class="go"> -->
          {{buttonMsg}}
        <!-- </div> -->
      </button>
      <h2>{{result}}</h2>
    </div>

  </div>
</template>

<script>
  export default{
    props:[
      'content',
    ],
    data(){
      return {
        result:"",
        input:"",
        buttonMsg:"keine Ahnung",
      }
    },
    created(){
      this.action=this.go
    },
    methods:{
      go(){
        if (this.input==""){

          return
        }
        if (this.input==this.content.solution){
          this.result=this.content.solution+" ist richtig!"
          this.input=""
          this.buttonMsg="nächste"
          this.$emit('solved')
          this.action=this.next
        }else{
          this.result=this.input+" ist leider Falsch. "+this.content.solution+' ist die richtige Antwort.'
          this.input=""
          this.buttonMsg="nächste"
          this.$emit('failed')
          this.action=this.next
        }
      },
      next(){
        this.$emit('next')
        this.buttonMsg="aufgeben"
        this.action=this.go
        this.input=""
      },
      keyup(k){
        console.log(k.key);
        if (k.key=="Enter"){
          this.action()
        }else{
          this.buttonMsg="fertig"
        }
      },

      start(){
        this.$refs.input.focus()
        this.result="..."
      }
    },
  }
</script>
