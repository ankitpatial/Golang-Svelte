## Job Flow

- Request and estimate. **Status: Created**
    - Send for measurement **Status: Estimating**
    - After measurement, price and move estimates. **Status: Estimated**
- Estimate:
    - Approve, will be now considered as new job. **Status: Approved**
    - Disapprove, mean client is not ready to go with job. **Status: Disapproved**
- Jobs that are ready to go, assign Partner As **Assigned** (send invite to partner)
  rule 01, partner has 2 days to accept or reject the job
  - **accepted**, on accepted, move to **InProgress**
  - **rejected**, on rejection move back to status **Approved**
- InProgress
- Done

