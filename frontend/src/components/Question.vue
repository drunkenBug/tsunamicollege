<template>
  <div class="bubble">
    <div class="question">

      <h2>{{title}}</h2>
      <h2>{{result}}</h2>

      <input v-if="showInput" ref="input" class="test" type="number" v-model='input' @keyup="keyup" >

      <button class="test" type="button" name="button" ref="action" @click='action' >
        <!-- <div class="go"> -->
          {{buttonMsg}}
        <!-- </div> -->
      </button>
    </div>

  </div>
</template>

<script>
  export default{
    props:[
      'content',
    ],
    watch:{
      content(newTitle){
        this.title=newTitle.title
      }
    },
    data(){
      return {
        title:this.content.title,
        showInput:true,
        result:"",
        input:"",
        buttonMsg:"keine Ahnung",
      }
    },
    created(){
      this.action=this.go
      this.$nextTick(()=>{
        this.$refs.input.focus()
        console.log("input found:",this.$refs.input);
      })
    },
    methods:{
      go(){
        if (this.input==""){

          return
        }
        if (this.input==this.content.solution){
          this.input=""
          var newglow=getComputedStyle(document.documentElement).getPropertyValue('--glow')
          newglow=parseFloat(newglow)+0.1
          document.documentElement.style.setProperty('--glow',newglow)
          this.$emit('solved')
          this.$emit('next')
          return
        }else{
          this.result=this.input+" ist leider Falsch. "+this.content.solution+' ist die richtige Antwort.'
          this.$emit('failed')
          this.input=""
          this.action=this.next
          this.buttonMsg="next"
          this.showInput=false
          this.$refs.action.focus()
        }


      },
      next(){

        this.$emit('next')
        this.buttonMsg="aufgeben"
        this.action=this.go
        this.input=""
        this.showInput=true
        // console.log(this.$refs.input.childNodes[0]);
        this.$nextTick(()=>{
          this.$refs.input.focus()
        })

      },
      keyup(k){
        if (k.key=="Enter"){
          this.action()
        }else{
          this.buttonMsg="fertig"
        }
      },

      start(){
        // this.$refs.input.focus()
        this.result=""

      }
    },
  }
</script>
