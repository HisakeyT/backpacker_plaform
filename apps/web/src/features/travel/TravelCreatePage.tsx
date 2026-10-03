import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Box, Typography } from "@mui/material";
import { useAuth } from "../../app/providers/useAuth.ts";
import { createTravel } from "./repository";
import { TravelForm } from "./TravelForm";
import type { CreateTravelInput } from "./types";

export const TravelCreatePage = () => {
  const { token } = useAuth();
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (input: CreateTravelInput) => {
    if (!token) return;
    setIsSubmitting(true);
    setError(null);
    try {
      const travel = await createTravel(token, input);
      navigate(`/travels/${travel.id}`);
    } catch {
      setError("旅行の作成に失敗しました");
      setIsSubmitting(false);
    }
  };

  return (
    <Box className="TravelCreatePage">
      <Typography
        variant="h4"
        component="h1"
        sx={{
          mb: 3,
          fontSize: { xs: "1.5rem", sm: "2.125rem" },
          wordBreak: "break-word",
        }}
      >
        新しい旅行を作る
      </Typography>
      <TravelForm
        submitLabel="作成する"
        error={error}
        isSubmitting={isSubmitting}
        onSubmit={handleSubmit}
      />
    </Box>
  );
};
