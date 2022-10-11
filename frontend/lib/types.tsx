export interface ILillyVideo {
	id: string;
	cid: string;
	account?: string;
	nickname: string;
	name: string;
	description: string;
	type: string;
	categories: string[];
	video_uri: string;
	thumbnail_uri: string;
	views: number;
	created_at: Date;
};

export interface IUserInfo {
  address?: string,
  nickname: string,
  login_request_id: string,
  videos?:string,
  comments?: string,
  vid_reports?: string,
  ads?: string,
  created_at?: string,
  updated_at?: string,
  ID?: string,
};
export interface IComment {
  id?: string;
  nickname: string;
  description: string;
  created_at?: string;
  UpdatedAt?: string;
};
export interface IInfinite {
  result: ILillyVideo[];
  current_page: number;
  isLast: boolean;
};
export interface IModal {
  title: string;
  description1: string;
  description2: string;
}
export interface IStakeRequest {
  amount: string;
  stakedBlock: string;
  status: string;
  unstakeRequestBlock: string;
};