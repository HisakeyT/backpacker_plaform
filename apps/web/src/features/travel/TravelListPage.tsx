import { useEffect, useState } from "react";
import { useAuth } from "../../app/providers/useAuth";
import { getTravels } from "./repository";
import type { Travel } from "./types";
import { Box, Typography } from "@mui/material";
import { TravelCard } from "./TravelCard";

export const TravelListPage = () => {
  const { token, user } = useAuth();
  const [travels, setTravels] = useState<Travel[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!token) return;

    const fetchTravels = async () => {
      try {
        const data = await getTravels(token);
        setTravels(data);
      } catch {
        setError("旅行の取得に失敗しました");
      } finally {
        setIsLoading(false);
      }
    };

    fetchTravels();
  }, [token]);

  if (isLoading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        {user?.nickname}の旅行記
      </Typography>

      {travels.length === 0 ? (
        <Typography color="text.secondary">
          まだ旅の記録がありません。最初の一歩を記録しよう
        </Typography>
      ) : (
        <Box
          sx={{
            display: "grid",
            gap: 2,
            gridTemplateColumns: "repeat(auto-fill, minmax(260px, 1fr))",
          }}
        >
          {travels.map((travel) => (
            <TravelCard key={travel.id} travel={travel} />
          ))}
        </Box>
      )}
    </Box>
  );
};
