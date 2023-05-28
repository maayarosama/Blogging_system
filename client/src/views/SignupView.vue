<template>
    <div class="div-wrapper">
        <v-container>
            <v-row justify="center">
                <v-col cols="12" sm="6">
                    <h5 class="text-h5 text-md-h4 font-weight-bold text-center mt-10 primary">
                        Create a new account </h5>

                    <v-form v-model="verify" @submit.prevent="onSubmit">

                        <BaseInput label="Name" :rules="nameRules" type="text" placeholder="Please enter your full name"
                            @emit-update-input="updateInput" />

                        <BaseInput label="Email" :rules="emailRules" type="text" placeholder="Please enter your email"
                            @emit-update-input="updateInput" />
                        <v-textarea v-model="quote" :rules="Rules" label="Quote"
                            placeholder="Please enter your favourite quote" bg-color="accent" variant="outlined"
                            class="my-2">
                        </v-textarea>
                        <BaseInput label="Password" :rules="passwordRules" type="password"
                            placeholder="Please enter your password" @emit-update-input="updateInput" />

                        <BaseInput label="Confrim Password" :rules="cpasswordRules" type="password"
                            placeholder="Please Re-enter your password" @emit-update-input="updateInput" />
                        <!-- <BaseButton type="submit" :disabled="!verify || alert" class="text-capitalize mx-auto my-5 bg-primary"
                    :loading="loading" text="Signin" variant="flat" /> -->

                        <v-btn type="submit" block :loading="loading" variant="flat" color="primary"
                            class="text-capitalize mx-auto my-5 bg-primary">
                            Login
                        </v-btn>
                        <div align="center">
                            <p>

                                <router-link class="text-body-2 text-decoration-none primary" to="/login">Back to
                                    login</router-link>
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
const quote = ref("");

const router = useRouter();

const updateInput = (value, label) => {

    inputValue.value.set(
        label, value
    )

}

const Rules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (value.length < 3) return "Field should be at least 3 characters";
        return true;
    },
]);

const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

const emailRules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (!value.match(emailRegex)) return "Invalid email address";
        return true;
    },
]);
const nameRegex = /^(\w+\s){0,3}\w*$/;

const nameRules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (!value.match(nameRegex)) return "Must be at most four names";
        if (value.length < 3) return "Field should be at least 3 characters";
        if (value.length > 20) return "Field should be at most 20 characters";
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

const cpasswordRules = ref([
    (value) => {
        if (!value) return "Field is required";
        if (value.length < 7) return "Password must be at least 7 characters";
        if (value.length > 12) return "Password must be at most 12 characters";
        if (value !== inputValue.value.get("Password")) return "Passwords don't match";
        return true;
    },
]);


const onSubmit = () => {
    if (!verify.value) return;

    loading.value = true;
    axios
        .post("http://localhost:5000/user/signup", {
            name: inputValue.value.get("Name"),
            email: inputValue.value.get("Email"),
            password: inputValue.value.get("Password"),
            quote: quote.value,
        })
        .then((response) => {
            localStorage.setItem("name", inputValue.value.get("Name"));
            localStorage.setItem("password", inputValue.value.get("Password"));
            localStorage.setItem("email", inputValue.value.get("Email"));
            localStorage.setItem("quote", quote.value);
            // toast.value.toast(response.data.msg);
            console.log(response)
            router.push({
                name: "OTP",
                query: {
                    email: inputValue.value.get("Email"),
                    // isSignup: isSignup.value,
                    // timeout: response.data.data.timeout,
                },
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