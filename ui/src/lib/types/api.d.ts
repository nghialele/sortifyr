export namespace API {
  interface Base extends JSON {
    id: number;
  }

  export interface User extends Base {
    uid: string;
    name: string;
    display_name: string;
    email: string;
  }

  export interface Playlist extends Base {
    spotify_id: string;
    owner?: User;
    name: string;
    description?: string;
    public: boolean;
    tracks: number;
    collaborative: boolean;
    updated_at: string;
  }
}
