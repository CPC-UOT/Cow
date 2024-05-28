<template>
  <div class="border"></div>
  <div class="background" dir="rtl">
    <img class="logo" src="../assets/logo.png" alt="" />
    <img class="nit" src="../assets/utsu.png" alt="" />
    <img class="nit" src="../assets/nit.png" alt="" />
    <img class="science" src="../assets/scienceLogo.png" alt="" />
  </div>
  <div class="container">
    <div class="farm-container">
      <Icon />
    </div>
    <div class="title" dir="rtl">
      <h1>سـاعد صاحب المزرعة</h1>
      <h1>في حلب أبقاره!</h1>
      <h1>&lt;/&gt; ++C</h1>
    </div>
  </div>
  <div class="details" dir="rtl">
    <div class="box-cotainer">
      <div class="box num">
        عدد المحاولات الناجحة
        <i class="fa-solid fa-check"></i>
      </div>
      <div class="box">
        {{ CorrectAnswers }}
        <i class="fa-solid fa-user"></i>
      </div>
    </div>
    <div class="box-cotainer">
      <div class="box red-box num">
        عدد المحاولات الفاشلة
        <i class="fa-solid fa-xmark"></i>
      </div>
      <div class="box red-box">
        {{ trying }}
        <i class="fa-solid fa-user"></i>
      </div>
    </div>
  </div>
  <div dir="rtl" class="info">
    <h1>
      <i class="fa-solid fa-certificate"></i>
      التفاصيل
    </h1>
    <h3>
      يفكر صاحب المزرعة في تغيير طريقة تخصيص الدلاء لحلب الأبقار, ويعتقد أن هذا
      سيسمح له في النهاية بإستخدام عدد اقل من إجمالي الدلاء ولكن لا يعرف كم
      بالضبط, ويحتاج مساعدتك!<br />
      لدى صاحب المزرعة عدد N من الأبقار
      <bdi>( 1 ≤ N ≤ 100 )</bdi>
      مُرقمة كالتالي 1...N, البقرة i ستحتاج أن يتم حلبها في زمن مُحدد من الوقت
      s<sub>i</sub> الى الوقت t<sub>i</sub> وستحتاج الى b<sub>i</sub> من الدلاء
      أثناء عملية الحلب.
      <br />
      صاحب المزرعة يريد طريقة لمعرفة العدد الأدنى من الدلاء لتخزينها حتى لا يخزن
      عدد أكبر من حاجته بحيث كل أبقاره يتم حلبهم في وقتهم المُناسب وبدون ايضاً
      ان يتفاجئ بعدم وجود دلاء تكفي لحلب جميع الأبقار.
    </h3>
    <h3>
      <span style="font-weight: 800">المدخلات</span>
      <br />
      سيكون لديك ملف blist.in يحتوي في أول سطر على عدد الأبقار N, وكل سطر بعد
      ذالك يحتوي على معلومات كل بقرة s<sub>i</sub> t<sub>i</sub> و b<sub>i</sub>
      مفصولين بمسافة. كلا من s<sub>i</sub> و t<sub>i</sub> عدد صحيح في نطاق
      1...1000 بينما b<sub>i</sub> عدد صحيح يتراوح بين 1...10.
    </h3>
    <h3>
      <span style="font-weight: 800">المخرجات</span>
      <br />
      سيكون لديك ملف blist.out المطلوب منك إدخال إليه العدد الأدنى من الدلاء
      التي يحتاجها صاحب المزرعة.
    </h3>
    <h3 style="">
      <span style="font-weight: 800">عينة توضيحية</span>
      <br />
      - المدخلات (من اليسار الى اليمين)
      <br />
      3
      <br />
      <bdi> 4 10 1</bdi>
      <br />
      <bdi> 8 13 3</bdi>
      <br />
      <bdi>2 6 2</bdi>
      <br />
      - المخرجات
      <br />
      4
      <br />
      في هاذه العينة أحتاج صاحب المزرعة الى 4 دلاء حيث سيحتاج الى 2  من دلاء لحلب البقرة من 2 الى 6 ولكن في الوقت بين ال4 و 6 سيحتاج صاحب الزرعة الى 3 من دلاء لانه هناك بقرة اخرى سيتم حلبها في نفس الوقت وعندما يدخل الوقت 8 سيحتاج الى 4 من دلاء وهاذا هوا الحد الأدنى المطلوب تخزينه.
    </h3>
  </div>
  <div class="editor">
    <monaco />
  </div>
  <div class="compiler">
    <div class="buttons-container">
      <button @click="runActived()" class="run">
        <i class="fa-solid fa-gear"></i>
        RUN
      </button>
      <button @click="sendActived()" class="run sent">
        <i class="fa-solid fa-paper-plane"></i> ارسل إجابتك
      </button>
    </div>
    <div class="messages-container">
      <div class="messages-error">
        <h1>
          <i class="fa-solid fa-circle-xmark"></i>
          الأخطاء
        </h1>
        <h2>{{ compErrString }}</h2>
      </div>
      <div class="messages-print">
        <h1><i class="fa-solid fa-print"></i> الطباعة</h1>
        <h2>{{ output }}</h2>
      </div>
    </div>
  </div>
</template>
<script setup>
import store from "../store/index.js";
import Monaco from "../components/Monaco.vue";
import Icon from "../components/Icon.vue";
</script>
<script>
import swal from "sweetalert2";
let urlPath = process.env.VUE_APP_BACKEND_URL
export default {
  name: "Home-view",
  data() {
    return {
      store,
      urlPath,
      trying: 0,
      CorrectAnswers: 0,
      checkAnswer: false,
      compErr: false,
      compErrString: "",
      output: "",
    };
  },
  mounted() {
    fetch(`${this.urlPath}/correct`)
      .then(res => (res.json()))
      .then(c => this.CorrectAnswers = c)

    fetch(`${this.urlPath}/wrong`)
      .then(res => res.json())
      .then(c => this.trying = c)
    this.setOutput("")
    this.setErrors("")
  },
  methods: {
    sendActived() {
      this.runActived()
      .then(() => {
      fetch(`${this.urlPath}/code/check`)
        .then(response => {
          console.log(response.status)
          if (response.status === 200) {
            return true
          } else {
            return false
          }
        })
        .then((res) => this.AnswerAlert(res))
      })
    },
    async runActived() {
      // كود لما ترص على run
      const requestOptions = {
          method: 'POST',
          headers: {
          },
          body: this.store.state.code
      };

      await fetch(`${this.urlPath}/code`, requestOptions)
        .then(response => {
          if (response.status === 400) {
            this.compErr = true;
          } else if (response.status === 200) {
            this.compErr = false;
          }
          return response
        })
        .then(response => response.json())
        //.then(j => JSON.parse(j))
        .then(output => {
          if (this.compErr === true) {
            this.setErrors(output)
            this.setOutput("")
          } else {
            this.setOutput(output)
            this.setErrors("")
          }
        })
      },
    setErrors(err) {
      if (this.compErr === true) {
        this.compErrString = err
        return
      }
      this.compErrString = "لايوجد...";
    },
    setOutput(out) {
      if (this.compErr === false && out != "") {
       this.output = out
        return
      }
      this.output = "لايوجد...";
    },
    AnswerAlert(res) {
      if (res === true) {
        swal
          .fire({
            html: `<dotlottie-player
          class="farm"
          src="https://lottie.host/0fa7adad-b236-4f41-92ab-58d91058dfdb/kmv95X11hf.lottie"
          background="transparent"
          speed="1"
          style="width: 500px; height: 400px"
          loop
          autoplay
        ></dotlottie-player>
        <h1>مبرووك! قمت بالمطلوب </h1>`,
            confirmButtonText: "حسناً",
          })
          .then(function () {
            window.location.reload();
          });
      } else {
        swal
          .fire({
            icon: "error",
            imageHeight: 300,
            imageAlt: "A short image",
            confirmButtonText: "حسناً",
            title: "إجابة خاطئة",
          })
          .then(function () {
            window.location.reload();
          });
      }
    },
  },
};
</script>
<style>
.the-container {
  height: 100vh;
  margin-bottom: 2rem;
}
.background {
  justify-content: center;
  display: flex;
  align-items: center;
  height: 8rem;
  width: 100%;
  background-image: linear-gradient(to right, #c0aef5, #e7f2ca, #caf2f2);
  margin: 0;
}
::selection {
  background-color: #caf2f2;
}
.title {
  color: #6835f8;
  margin-right: 12rem;
  margin-top: -0.1rem;
  border-radius: 1rem;
}
.title h1 {
  font-family: "Bahij TheSansArabic";
  font-weight: 800;
  font-size: 3rem;
  margin-bottom: -2.5rem;
}
.container {
  display: flex;
  justify-content: space-between;
  position: relative;
  margin-bottom: 4rem;
  margin-top: 5rem;
}
.farm-container {
  display: flex;
  margin-left: 4.8rem;
}
.logo {
  height: 50px;
  margin-right: 12rem;
  margin-left: 2rem;
}
.nit {
  height: 80px;
  margin-left: 2rem;
}
.science {
  height: 80px;
  margin-left: 12rem;
}
.icons img {
  height: 100px;
}
::-webkit-scrollbar {
  display: none;
}
.details {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding-top: 2rem;
  padding-bottom: 1rem;
  margin-right: 12rem;
}
.box {
  color: white;
  background-color: rgb(93, 189, 133);
  width: fit-content;
  min-width: 3rem;
  display: flex;
  justify-content: center;
  padding: 0.5rem 1rem;
  border-radius: 1.5rem;
  margin-left: 1rem;
}
.num {
  width: 12rem;
}
.red-box {
  background-color: rgb(0, 0, 0);
}
.fa-solid {
  padding: 0;
  margin: 0;
  margin-top: 0.3rem;
  margin-right: 0.4rem;
}
.fa-check {
  margin-right: 0.6rem;
}
.box-cotainer {
  display: flex;
  margin-bottom: 1rem;
}
.editor {
  z-index: 1;
  display: flex;
  justify-content: center;
}
.info {
  margin-right: 12rem;
}
.fa-certificate {
  margin-left: 1rem;
  color: #6835f8;
}
.info h3 {
  font-weight: 500;
  margin-right: 4rem;
  width: 80%;
}
.compiler {
  margin-left: 13rem;
  margin-top: -3rem;
  margin-bottom: 3rem;
}
.buttons-container {
  display: flex;
  margin-right: 12rem;
}
.run {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 1.5rem;
  color: white;
  background-color: rgb(93, 189, 133);
  padding-right: 1rem;
  padding-left: 1rem;
  border-radius: 1rem;
  margin-bottom: 1rem;
}
.run i {
  font-size: 1.2rem;
  margin-bottom: 0.3rem;
}
.sent {
  background: -webkit-linear-gradient(30deg, #6c17a0, #7d24eb);
  width: 12rem;
  margin-left: 1rem;
}
.messages-container {
  margin-right: 12rem;
  display: flex;
}
.messages-error {
  font-size: 0.8rem;
  color: rgb(152, 12, 45);
  height: 15rem;
  overflow-x: scroll;
  width: 60%;
  border: rgb(152 12 45/0.1) 5px solid;
  background-color: rgb(236, 135, 135);
  border-radius: 2rem;
  padding: 0.2rem 2rem;
}
.messages-error h2 {
  margin-left: 1rem;
  margin-top: -0.5rem;
}
.messages-print {
  margin-left: 1%;
  width: 39%;
  font-size: 0.8rem;
  color: #6835f8;
  height: 15rem;
  overflow-x: scroll;
  border: rgb(104 53 248/0.2) 5px solid;
  background-color: #b39cf3;
  border-radius: 2rem;
  padding: 0.2rem 2rem;
}
</style>
