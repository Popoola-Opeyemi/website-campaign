<template>
  <div>
    <nav-bar @register="register" />
    <section class="hero is-fullheight">
      <h1 class="title has-text-white">Affordable websites at low cost.</h1>
    </section>
    <b-loading :is-full-page="isFullPage" :active.sync="isLoading" :can-cancel="true"></b-loading>
    <section class="clients">
      <div class="container">
        <div class="section-title has-text-centered">Trusted by us</div>
        <div class="columns is-multiline">
          <div class="column has-text-centered is-3" v-for="icon in imgBox" :key="icon.key">
            <div class="img-box">
              <img :src="'/img/logo/'+icon.icon" alt />
            </div>
          </div>
        </div>
        <div class="title-offering has-text-centered">website design packages</div>
        <div class="title-content has-text-centered">
          <p>select the best website packages that match</p>
          <p>your requirements</p>
        </div>
      </div>
    </section>
    <section class="pricing-cards">
      <div class="custom-container">
        <div class="pricing">
          <div class="columns">
            <div class="column" v-for="(arr, idx) in Plans" :key="idx">
              <div class="box pricing-box">
                <div class="box-title has-text-centered">
                  <p class="main-title">{{arr.name}}</p>
                  <p class="middle">{{arr.details}}</p>
                  <p class="pg-count">{{arr.page_no}}</p>
                </div>
                <div class="prices">
                  <ul class="prices-list">
                    <li class="text">As low as</li>
                    <li class="currency">&#8358; {{arr.amount}}</li>
                    <li class="currency-strike">
                      <strike>&#8358; {{arr.discount}}</strike>
                    </li>
                  </ul>
                  <b-button
                    type="is-primary button-price"
                    @click="register(arr.name)"
                    expanded
                  >order</b-button>
                  <span class="check-icon">
                    <ul>
                      <li v-for="plans in arr.content" :key="plans.id">
                        <i class="mdi mdi-check"></i>
                        <i class="names">{{ plans.name }}</i>
                      </li>
                    </ul>
                  </span>
                </div>
              </div>
            </div>

            <!-- <div class="column">
              <div class="box pricing-box">Pricing</div>
            </div>
            <div class="column">
              <div class="box pricing-box">Pricing</div>
            </div>-->
          </div>
        </div>
      </div>
    </section>
    <section class="contact-pg">
      <div class="cus-contain">
        <div class="columns">
          <div class="column is-3">
            <div class="contact-box">
              <p class="_title">LET US HELP CREATE THAT AWESOME WEBSITE</p>
              <P class="_subtitle">Request a free Quote</P>
              <b-field
                label="Full Name"
                :type="errors.has('Full Name') ? 'is-danger' : ''"
                :message="
                  errors.has('Full Name') ? errors.first('Full Name') : ''
                "
              >
                <b-input
                  autocomplete="name"
                  name="Full Name"
                  v-validate="'required'"
                  placeholder="Full Name"
                  v-model="quote.fullname"
                />
              </b-field>

              <b-field
                label="Email"
                :type="errors.has('Email') ? 'is-danger' : ''"
                :message="errors.has('Email') ? errors.first('Email') : ''"
              >
                <b-input
                  name="Email"
                  v-validate="'required|email'"
                  placeholder="Email"
                  v-model="quote.email"
                />
              </b-field>

              <button class="button" @click="sendQuote">Send</button>
            </div>
          </div>
          <div class="column direct">
            <div class="about-content">
              <p class="_title">About Us</p>
              <div class="_content">
                <p>
                  Prudent Digital UK, a member of the Prudent Business Group
                  Limited is an IT solutions and services “One-stop shop” with
                  renowned specialist workforce and partners. Our expert
                  developers and professional IT team who have a staunch passion
                  for what we do together, have over 100 years of experience in
                  IT, eCommerce, Software Engineering, Creative Digital Designs
                  and General IT consultancy.
                </p>
                <p>
                  At Prudent Digital, we operate a person or business centred
                  approach based on individual needs and business requirements
                  with our expert team ready to help create e􀇖ective and
                  captivating event designs, easy to use websites and mobile
                  applications channelling your core messages to your desired
                  audience. We also o􀇖er ranges of Digital and IT training for
                  individuals and organizations
                </p>
                <p>
                  We pride ourselves in utmost professionalism and excellent
                  quality of product and services putting the most relevant
                  element in mind – COST. We o􀇖er a second- to-none cost
                  effective services and packages leaving our clients with a
                  smile and a little left for the rainy day.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="footer">
      <div class="mobile-footer">
        <div class="columns has-text-cenetered">
          <div class="column">
            <div class="logo-col">
              <a target="_blank" href="http://pbgdigital.co.uk">
                <img src="/img/website-logo.png" alt="logo" />
              </a>
            </div>
          </div>
          <div class="column">
            <div class="icon-set">
              <div class="icon-bg">
                <a target="_blank" href="https://twitter.com/DigitalPrudent">
                  <i class="mdi mdi-twitter"></i>
                </a>
              </div>
              <div class="icon-bg">
                <a target="_blank" href="https://www.facebook.com/pbgroupdigital/">
                  <i class="mdi mdi-facebook"></i>
                </a>
              </div>
              <div class="icon-bg">
                <a
                  target="_blank"
                  href="https://www.linkedin.com/company/prudent-business-group-ltd/?viewAsMember=true"
                >
                  <i class="mdi mdi-linkedin"></i>
                </a>
              </div>
              <div class="icon-bg">
                <a target="_blank" href="https://wa.me/447440713708">
                  <i class="mdi mdi-whatsapp"></i>
                </a>
              </div>
              <div class="icon-bg">
                <a target="_blank" href="https://www.instagram.com/prudent_digital/">
                  <i class="mdi mdi-instagram"></i>
                </a>
              </div>
            </div>
          </div>
          <div class="column">
            <a
              class="home-link"
              target="_blank"
              href="http://pbgdigital.co.uk"
            >www.pbgdigigtal.co.uk</a>
          </div>
        </div>
      </div>
    </section>
    <!-- <section class="spacing"></section> -->
  </div>
</template>

<script>
import NavBar from "@/components/navbar";
import RegisterForm from "@/components/RegisterForm";
import { Notification } from "@/utils/helpers";

export default {
  layout: "website",
  components: {
    NavBar,
    RegisterForm
  },
  data() {
    return {
      quote: {},
      isFullPage: true,
      isLoading: false,
      imgBox: [
        { icon: "Logo-1.png" },
        { icon: "Logo-2.png" },
        { icon: "Logo-3.png" },
        { icon: "Logo-4.png" },
        { icon: "Logo-5.png" },
        { icon: "Logo-6.png" },
        { icon: "Logo-7.png" },
        { icon: "Logo-8.png" },
        { icon: "Logo-9.png" },
        { icon: "Logo-10.png" },
        { icon: "Logo-11.png" },
        { icon: "Logo-12.png" }
      ],
      Plans: [
        {
          name: "Starter",
          amount: "149,599",
          details: "Static customized website",
          page_no: "3 pages",
          discount: "199,599",
          content: [
            { name: "unlimited Revision for Homepage Design" },
            { name: "Responsive" },
            { name: "Contact Form" },
            { name: "ALways Faster" },
            { name: "Secured" },
            { name: "Organized" },
            { name: "Money Back Guarantee" },
            { name: "Web Hosting" }
          ]
        },
        {
          name: "Growth",
          amount: "249,599",
          details: "website with content managment system ",
          page_no: "6 pages",
          discount: "299,599",
          content: [
            { id: 1, name: "unlimited Revision for Homepage Design" },
            { id: 2, name: "Responsive" },
            { id: 3, name: "Contact Form" },
            { id: 4, name: "SEO" },
            { id: 6, name: "ALways Faster" },
            { id: 7, name: "Secured" },
            { id: 8, name: "Organized" },
            { id: 9, name: "Custom Theme" },
            { id: 11, name: "Money Back Guarantee" },
            { id: 13, name: "Web Hosting" }
          ]
        },
        {
          name: "Advanced",
          amount: "349,599",
          details:
            "website with content managment system (cms) and custom pages ",
          page_no: "12 pages",
          discount: "399,599",
          content: [
            { name: "unlimited Revision for Homepage Design" },
            { name: "Responsive" },
            { name: "Contact Form" },
            { name: "Technical SEO" },
            { name: "ALways Faster" },
            { name: "Secured" },
            { name: "Organized" },
            { name: "Custom Theme" },
            { name: "Custom Email Templating" },
            { name: "Socail Media" },
            { name: "Money Back Guarantee" },
            { name: "Header SlideShow" },
            { name: "Web Hosting" }
          ]
        }
      ]
    };
  },
  methods: {
    register(plan) {
      let that = this;
      this.$buefy.modal.open({
        parent: this,
        component: RegisterForm,
        hasModalCard: true,
        width: "900",
        customClass: "register-form",
        events: {
          save(e) {
            that.sendForm(e, plan);
          }
        }
      });
    },
    async sendQuote() {
      let formdata = {
        url: "quote",
        data: this.quote
      };

      let result = await this.$validator.validateAll();
      if (result) {
        this.isLoading = true;
        try {
          let req = await this.$store.dispatch("Postdata", formdata);
          if (req) {
            this.isLoading = false;

            Notification(
              this,
              "Thank you for filling the form",
              "is-success",
              "is-top",
              1000
            );
          } else {
            this.isLoading = false;

            Notification(
              this,
              "An error occured. Please try again ",
              "is-danger",
              "is-top",
              1000
            );
          }
        } catch (error) {
          this.isLoading = false;
          Notification(
            this,
            "An error occured. Please try again ",
            "is-danger",
            "is-top",
            1000
          );
        }
      }
    },

    async sendForm(form, plan) {
      if (plan != null) {
        form.plan_type = plan;
      }
      let formdata = {
        url: "register",
        data: form
      };
      this.isLoading = true;

      try {
        let req = await this.$store.dispatch("Postdata", formdata);
        if (req) {
          this.isLoading = false;
          Notification(
            this,
            "Thank you for filling the form",
            "is-success",
            "is-top",
            1000
          );
        } else {
          this.isLoading = false;
          Notification(
            this,
            "An error occured. Please try again ",
            "is-danger",
            "is-top",
            1000
          );
        }
      } catch (error) {
        this.isLoading = false;

        Notification(
          this,
          "An error occured. Please try again ",
          "is-danger",
          "is-top",
          1000
        );
      }
    }
  }
};
</script>

<style></style>
