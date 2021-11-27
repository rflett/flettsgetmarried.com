import {IGuest, IGuestsResult, ISearchMatch, ISearchResult} from "~/models/interfaces";

export class Guest implements IGuest {
  id: string;
  inviteId: string;
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

  constructor(fromServer: IGuest) {
    this.id = fromServer.id;
    this.inviteId = fromServer.inviteId;
    this.firstName = fromServer.firstName;
    this.lastName = fromServer.lastName;
    this.rsvp = fromServer.rsvp;
    this.vaccinated = fromServer.vaccinated;
    this.createdAt = fromServer.createdAt;
    this.updatedAt = fromServer.updatedAt;
    this.mobile = fromServer.mobile;
    this.email = fromServer.email;
    this.diet = fromServer.diet;
    this.music = fromServer.music;
  }
}

export class SearchMatch implements ISearchMatch {
  guests: IGuest[];
  inviteId: string;

  constructor(fromServer?: ISearchMatch) {
    if (fromServer) {
      this.guests = fromServer.guests;
      this.inviteId = fromServer.inviteId;
    } else {
      this.guests = [];
      this.inviteId = "";
    }
  }

  toDto(): ISearchMatch {
    return {
      inviteId: this.inviteId,
      guests: this.guests
    }
  }
}

export class SearchResult implements ISearchResult {
  matches: SearchMatch[];
  partialMatch: boolean;

  constructor(fromServer?: ISearchResult) {
    if (fromServer) {
      this.matches = fromServer.matches.map(match => new SearchMatch(match));
      this.partialMatch = fromServer.partialMatch;
    } else {
      this.matches = [];
      this.partialMatch = true;
    }
  }
}

export class GuestsResult implements IGuestsResult {
  guests: IGuest[];

  constructor(fromServer?: IGuestsResult) {
    if (fromServer) {
      this.guests = fromServer.guests
    } else {
      this.guests = [];
    }
  }
}

