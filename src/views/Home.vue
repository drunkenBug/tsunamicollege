<template>
  <div class="home">

    <div class="upper">

      <div class="face">
        <img alt="Bq logo" src="../assets/Tc.png">
      </div>

      <div class="horizontal"></div>
    </div>

    <div class="motto">
      <p>Don't be better than them</p>
      <p>Be better than yesterday.</p>
    </div>

    <div class="entrance">
      <button type="button" name="button">
        <router-link to="/Test">Take the entrance test</router-link>
      </button>
    </div>


  </div>
</template>

<script>

export default {
  name: 'Home',
  components: {
  }
}
</script>

<style media="screen">
.upper{
  background: white;
  border-radius: 2rem;
}
.face{
  padding:1rem 0 2rem 0;
  border:1rem;
  border-radius:2rem;
}

.horizontal{
  background: #00f;
  padding: 2rem 0 0 0;
  border-radius: 0 0 2rem 2rem;
  margin: 0;
  border: 0;
  width: 100%;
  left: 0;
}

.motto{
  color: #00a;
  font-size: 20px;
  padding: 0;
  margin: 0;
  font-weight: bold;
}

.entrance button{
  background: #00f;
  border: 0;
  border-radius: 20px;
  color: white;
  margin-top: 20px;
  padding: 10px 20px;
}
.entrance a{
  color:white;
}
</style>
