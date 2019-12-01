export const state = () => ({
  Navigations: [
    {
      id: 1,
      name: "Home",
      link: "/"
    },
    {
      id: 2,
      name: "About",
      link: "/about"
    },
    {
      id: 3,
      name: "Our Works",
      link: "/works"
    },
    {
      id: 4,
      name: "Blog",
      link: "/blog"
    },
    {
      id: 5,
      name: "Team",
      link: "/team"
    },
    {
      id: 6,
      name: "Careers",
      link: "/careers"
    },
    {
      id: 7,
      name: "Contact",
      link: "/contact"
    }
  ]
});

export const mutations = {
  setbanner(state, data) {
    state.Banner = data;
  },
  setdisplayCaption(state, data) {
    state.displayCaption = data;
  },
  setcarousel(state, data) {
    state.Carousel = data;
  },

  savecontact(state, data) {
    state.contact = data;
  }
};

export const actions = {
  async Postdata({ commit }, param) {
    try {
      let req = await this.$axios.$post(`api/${param.url}`, param.data);
      return req.done;
    } catch (error) {
      throw error;
      console.log("an error occured", error);
    }
  }
};
