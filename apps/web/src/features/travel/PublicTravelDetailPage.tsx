import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { Box, Button, Typography } from "@mui/material";
import { PublicTravelPlanList } from "../travelPlan/PublicTravelPlanList";
import { getPublicTravel } from "./repository";
import type { PublicTravel } from "./types";

export const PublicTravelDetailPage = () => {
  const { travelId } = useParams();
  const [travel, setTravel] = useState<PublicTravel | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!travelId) return;

    const fetchTravel = async () => {
      try {
        const data = await getPublicTravel(Number(travelId));
        setTravel(data);
      } catch {
        setError("旅行が見つかりませんでした");
      } finally {
        setIsLoading(false);
      }
    };

    fetchTravel();
  }, [travelId]);

  if (isLoading) return <Typography>Loading...</Typography>;
  if (error || !travel) {
    return (
      <Typography color="error">
        {error ?? "旅行が見つかりませんでした"}
      </Typography>
    );
  }

  return (
    <Box className="PublicTravelDetailPage" sx={{ p: 2 }}>
      <Button component={Link} to="/public/travels" sx={{ mb: 2 }}>
        ← みんなの旅行記へ
      </Button>

      <Typography
        variant="h4"
        component="h1"
        sx={{
          fontSize: { xs: "1.5rem", sm: "2.125rem" },
          wordBreak: "break-word",
          mb: 1,
        }}
      >
        {travel.title}
      </Typography>

      <Typography color="text.secondary" gutterBottom>
        {travel.startDate} 〜 {travel.endDate}
      </Typography>
      <Typography variant="body2" color="text.secondary">
        by {travel.authorNickname}
      </Typography>

      <Typography variant="h5" component="h2" sx={{ mt: 4, mb: 2 }}>
        旅のプラン
      </Typography>
      <PublicTravelPlanList travelId={travel.id} />
    </Box>
  );
};
