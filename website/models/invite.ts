export interface Guest {
  id: string;
  inviteID: string;
  firstName: string;
  lastName: string;
  rsvp: boolean;
  vaccinated: boolean;
  createdAt: string;
  updatedAt: string | null;
  mobile: string | null;
  email: string | null;
  diet: string | null;
  music: string | null;
}

export interface SearchMatch {
  inviteID: string;
  guests: Guest[];
}

export interface SearchResult {
  partialMatch: boolean;
  matches: SearchMatch[];
}

export const GuestKara: Guest = {
  id: "aa1c37dd-8dc2-4338-a9c6-c73a72b61b2c",
  inviteID: "1",
  firstName: "kara",
  lastName: "hazelman",
  rsvp: false,
  vaccinated: false,
  createdAt: "2021-11-10T08:10:28Z",
  updatedAt: null,
  mobile: null,
  email: "kara.hazelman@hotmail.com",
  diet: null,
  music: null,
}

export const GuestRyan: Guest = {
  id: "d057f725-1d69-41d5-b69b-5c449cfdd0b7",
  inviteID: "1",
  firstName: "ryan",
  lastName: "flett",
  rsvp: false,
  vaccinated: false,
  createdAt: "2021-11-10T08:10:28Z",
  updatedAt: null,
  mobile: "0429235371",
  email: null,
  diet: null,
  music: null,
}

export const GuestMitch: Guest = {
  id: "65b6a783-d516-4c5c-8215-f799993b9a08",
  inviteID: "2",
  firstName: "mitch",
  lastName: "tinning",
  rsvp: false,
  vaccinated: false,
  createdAt: "2021-11-10T08:10:28Z",
  updatedAt: null,
  mobile: null,
  email: null,
  diet: null,
  music: null,
}

export const GuestNicole: Guest = {
  id: "6e2180fe-d040-4916-9aa4-b40bbdc8917c",
  inviteID: "2",
  firstName: "nicole",
  lastName: "dickinson",
  rsvp: false,
  vaccinated: false,
  createdAt: "2021-11-10T08:10:28Z",
  updatedAt: null,
  mobile: null,
  email: null,
  diet: null,
  music: null,
}

export const TestMatchOne: SearchMatch = {
  inviteID: "1",
  guests: [
    GuestRyan,
    GuestKara,
  ]
}

export const TestMatchTwo: SearchMatch = {
  inviteID: "2",
  guests: [
    GuestMitch,
    GuestNicole,
  ]
}

export const TestData: SearchResult = {
  partialMatch: false,
  matches: [
    TestMatchOne,
    TestMatchTwo,
  ]
}
