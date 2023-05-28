<template>
    <div class="div-wrapper">
        <v-container>
            <v-row justify="center">
                <v-col>
        <h5
          class="text-h5 text-md-h4 font-weight-bold text-center mt-10 primary"
        >
          Welcome To Blogging System
        </h5>
        <p class="text-center mb-10">
          Sign in to continue
        </p>
            <v-form v-model="verify" @submit.prevent="onSubmit">


                <BaseInput label="Email" :rules="emailRules" type="text" placeholder="Please enter your email"
                    @emit-update-input="updateInput" />

                <BaseInput label="Password" :rules="passwordRules" type="password" placeholder="Please enter your password"
                    @emit-update-input="updateInput" />
                <!-- <BaseButton type="submit" :disabled="!verify || alert" class="text-capitalize mx-auto my-5 bg-primary"
                    :loading="loading" text="Signin" variant="flat" /> -->

                    <v-btn
            type="submit"
            block
            :loading="loading"
            variant="flat"
            color="primary"
            class="text-capitalize mx-auto my-5 bg-primary"
          >
            Login
          </v-btn>
          <div align="center">
            <p>
              Don't have an account?
              <router-link
                class="text-body-2 text-decoration-none primary"
                to="/signup"
              >Sign up</router-link>
            </p>

          </div>
     
            </v-form>
                </v-col>

        </v-row>
        </v-container>

    </div>
</template>

<script setup>
// import BaseButton from "@/components/BaseButton.vue"

import BaseInput from "@/components/BaseInput.vue"
import { ref } from "vue"
import axios from "axios";
import { useRouter } from "vue-router";


const inputValue = ref(new Map());
const verify = ref(false);
const loading = ref(false);
const router = useRouter();

const updateInput = (value, label) => {

inputValue.value.set(
    label, value
)
// console.log(inputValue.value)
// inputValue.value = {label,value}
}

const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

const emailRules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (!value.match(emailRegex)) return "Invalid email address";
        return true;
    },
]);

const passwordRules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (value.length < 7) return "Password must be at least 7 characters";
        // if (value.length > 12) return "Password must be at most 12 characters";
        return true;
    },
]);


const onSubmit = () => {
      if (!verify.value) return;

      loading.value = true;
      axios
        .post("http://localhost:5000/user/signin", {
          email: inputValue.value.get("Email"),
          password: inputValue.value.get("Password"),
        })
        .then((response) => {
            console.log(response)
        //   localStorage.setItem("token", response.data.data.access_token);
        //   toast.value.toast(response.data.msg);
          router.push({
            name: "Home",
          });
        })
        .catch((error) => {
        //   toast.value.toast(error.response.data.err, "#FF5252");
        console.log(error)
          loading.value = false;
        });
    };

</script>

<style lang="scss" scoped>
.v-text-field .v-input__control .v-input__slot {
    min-height: auto !important;
    display: flex !important;
    align-items: center !important;
}
</style>