import { Card, CardContent, Typography } from "@mui/material";
import type { TravelPlan } from "./types";

type Props = {
  plan: TravelPlan;
};

export const TravelPlanCard = ({ plan }: Props) => {
  return (
    <Card className="TravelPlanCard" sx={{ height: "100%" }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          {plan.place}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {plan.content}
        </Typography>
      </CardContent>
    </Card>
  );
};
