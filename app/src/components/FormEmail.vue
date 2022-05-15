<template>
  <div
    class="formEmail w-full h-full flex flex-col justify-center items-center rounded-lg p-10 bg-white lg:bg-transparent"
  >
    <AlertMessage ref="alert" />
    <h1 class="logo text-3xl text-yellow-500">OpenMail</h1>
    <!-- <span class="text-yellow-500"
      ><i>Envie mensagens de um jeito rápido e fácil</i></span
    > -->
    <!-- <h4 class="text-lg text-yellow-500"><i>Rápido e fácil</i></h4> -->

    <form class="w-full flex flex-col mt-10">
      <input
        v-model="subject"
        type="text"
        placeholder="Assunto"
        :class="{ error: subjectNotValid }"
      />
      <input
        v-model="email"
        type="email"
        placeholder="Destinatário"
        :class="{ error: emailNotValid }"
      />
      <textarea
        v-model="message"
        placeholder="Mensagem"
        rows="5"
        :class="{ error: messageNotValid }"
      ></textarea>

      <button
        @click.prevent="send"
        :disabled="emailNotValid || messageNotValid || inProcess"
      >
        Enviar
      </button>
    </form>
  </div>
</template>

<script>
import { ref, computed } from "vue";
import AlertMessage from "@/components/AlertMessage";

export default {
  components: {
    AlertMessage,
  },
  setup() {
    // DATA
    const alert = ref(null);
    const subject = ref("Aqui vai um assunto legal");
    const email = ref("remetente@gmail.com");
    const message = ref("Aqui vai a mensagem que você deseja enviar");
    const inProcess = ref(false);

    // COMPUTED
    const subjectNotValid = computed(() => subject.value.trim() === "");
    const emailNotValid = computed(() => !/\S+@\S+\.\S+/.test(email.value));
    const messageNotValid = computed(() => message.value.trim() === "");
    const urlFetch = computed(
      () =>
        `${getHost()}/${encodeURIComponent(subject.value)}/${encodeURIComponent(email.value)}/${encodeURIComponent(
          message.value
        )}`
    );

    // METHODS
    const getHost = () => {
      const parts = window.origin.split(":");
      parts.pop();
      return `${parts.join(":")}:3000`;
    };
    const request = async () => {
      const { ok } = await fetch(urlFetch.value);

      if (ok) return success();

      fail();
    };
    const send = async () => {
      inProcess.value = true;
      try {
        await request();
      } catch {
        fail();
      }
      inProcess.value = false;
    };

    const success = () => alert.value.show(true, "Email enviado com sucesso");
    const fail = () => alert.value.show(false, "Erro ao enviar o email");

    return {
      alert,
      subject,
      email,
      message,
      inProcess,
      subjectNotValid,
      emailNotValid,
      messageNotValid,
      send,
    };
  },
};
</script>

<style lang="scss" scoped>
.logo {
  font-family: "Lobster", cursive;
}

input,
textarea {
  width: 100%;

  border: 2px solid #f5f5f5;
  border-radius: 7px;

  outline: none;

  margin-bottom: 15px;
  padding: 10px 15px;

  &.error {
    border-color: #e30e0e !important;
  }
  &:focus {
    border: 2px solid #eab308;
  }
}
button {
  color: white;
  font-weight: 500;
  background: #eab308;
  border-radius: 7px;
  padding: 10px 15px;

  &:disabled {
    background: #eee !important;
    pointer-events: none;
  }
}

.formEmail {
  max-width: 400px;
}
</style>
