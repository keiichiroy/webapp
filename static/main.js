var app = new Vue({
  el: '#app',
  data: {
    members: [],
  },
  mounted() {
    this.name = "Thiago"
    axios.get('/members').then(response => {
      this.members = response.data.members
    }
    ).catch(error => {
      console.log(error)
    })
  }
})