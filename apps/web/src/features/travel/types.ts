export type Travel = {
  id: number;
  userId: number;
  title: string;
  startDate: string;
  endDate: string;
  isPublic: boolean;
  createdAt: string;
  updatedAt: string;
}

export type CreateTravelInput = {
  title: string;
  startDate: string; // YYYY-MM-DD
  endDate: string;
  isPublic: boolean;
}

