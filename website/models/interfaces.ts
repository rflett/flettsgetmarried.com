export interface IGuest {
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
}

export interface ISearchMatch {
  inviteId: string;
  guests: IGuest[];
}

export interface ISearchResult {
  partialMatch: boolean;
  matches: ISearchMatch[];
}

export interface IGuestsResult {
  guests: IGuest[];
}
