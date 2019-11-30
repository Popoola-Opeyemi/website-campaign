<template>
  <div>
    <section class="hero is-fullheight">
      <h1 class="title has-text-white">Affordable websites at low cost.</h1>
      <!-- <div class="hero-body">
        <div class="container"></div>
      </div>-->
    </section>
    <section class="clients">
      <div class="container">
        <div class="section-title has-text-centered">Trusted by us</div>
        <div class="columns is-multiline">
          <div
            class="column has-text-centered is-3"
            v-for="icon in imgBox"
            :key="icon.key"
          >
            <div class="img-box">
              <img :src="'logo/' + icon.icon" alt />
            </div>
          </div>
        </div>
        <div class="title-offering has-text-centered">
          website design packages
        </div>
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
            <div class="column" v-for="a in 3" :key="a.id">
              <div class="box pricing-box">
                <div class="box-title has-text-centered">
                  <p class="main-title">Starter</p>
                  <p class="middle">Static Customized website</p>
                  <p class="pg-count">3 pages</p>
                </div>
                <div class="prices">
                  <ul class="prices-list">
                    <li class="text">As low as</li>
                    <li class="currency">&#8358; 149,599</li>
                    <li class="currency-strike">
                      <strike>&#8358; 149,599</strike>
                    </li>
                  </ul>
                  <b-button type="is-primary button-price" expanded
                    >Buy Now</b-button
                  >
                  <span class="check-icon">
                    <ul>
                      <li v-for="plans in Plans.starter" :key="plans.id">
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
                  v-model="form.name"
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
                  v-model="form.email"
                />
              </b-field>

              <button class="button" @click="send">Send</button>
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
    <footer class="footer">
      <div class="columns has-text-cenetered">
        <div class="column">
          <div class="logo-col">
            <a target="_blank" href="http://pbgdigital.co.uk">
              <img src="logo.png" alt="logo" />
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
              <a
                target="_blank"
                href="https://www.facebook.com/pbgroupdigital/"
              >
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
              <a
                target="_blank"
                href="https://www.instagram.com/prudent_digital/"
              >
                <i class="mdi mdi-instagram"></i>
              </a>
            </div>
          </div>
        </div>
        <div class="column">
          <a class="home-link" target="_blank" href="http://pbgdigital.co.uk"
            >www.pbgdigigtal.co.uk</a
          >
        </div>
      </div>
    </footer>
    <!-- <section class="spacing"></section> -->
  </div>
</template>

<script>
export default {
  layout: "website",
  data() {
    return {
      form: {},
      Plans: {
        starter: [
          { id: 1, name: "unlimited Revision for Homepage Design" },
          { id: 2, name: "Responsive" },
          { id: 3, name: "Contact Form" },
          { id: 4, name: "Technical SEO" },
          { id: 6, name: "ALways Faster" },
          { id: 7, name: "Secured" },
          { id: 8, name: "Organized" },
          { id: 9, name: "Custom Theme" },
          { id: 10, name: "Socail Media" },
          { id: 11, name: "Money Back Guarantee" },
          { id: 12, name: "Header SlideShow" },
          { id: 13, name: "Web Hosting" }
        ],
        Growth: [
          { id: 1, name: "unlimited Revision for Homepage Design" },
          { id: 2, name: "Responsive" },
          { id: 3, name: "Contact Form" },
          { id: 4, name: "Technical SEO" },
          { id: 6, name: "ALways Faster" },
          { id: 7, name: "Secured" },
          { id: 8, name: "Organized" },
          { id: 9, name: "Custom Theme" },
          { id: 10, name: "Socail Media" },
          { id: 11, name: "Money Back Guarantee" },
          { id: 12, name: "Header SlideShow" },
          { id: 13, name: "Web Hosting" }
        ]
      },
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
      ]
    };
  },
  methods: {
    async send() {
      let result = await this.$validator.validateAll();

      if (!result) {
        return;
      }
      try {
        let contact = await this.$store.dispatch("Postdata", param);

        if (!_.isEmpty(contact)) {
          Notification(this, `contact form sent suceessfully`, "is-top-right");
          setTimeout(() => window.location.reload(true), 600);
          this.$router.push("/contact");
        }
      } catch (error) {
        Notification(this, `an error occured`, "is-danger");
      }
    }
  }
};
</script>

<style></style>
