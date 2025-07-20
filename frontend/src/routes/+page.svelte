<script>
  import { onMount } from 'svelte';
  import MealInput from '$lib/MealInput.svelte';
  import { getMealPlans } from '$lib/api.js';
  
  const days = ['MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'];
  const dayLabels = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  const meals = ['BREAKFAST', 'LUNCH', 'DINNER'];
  const mealLabels = ['Breakfast', 'Lunch', 'Dinner'];
  
  let mealPlans = {};
  let loaded = false;
  
  onMount(async () => {
    try {
      const plans = await getMealPlans();
      const newMealPlans = {};
      plans.forEach(plan => {
        const key = `${plan.dayOfWeek}-${plan.mealType}`;
        newMealPlans[key] = plan.meal.name;
      });
      mealPlans = newMealPlans;
      loaded = true;
    } catch (error) {
      console.error('Error loading meal plans:', error);
      loaded = true;
    }
  });
  
  function getMealValue(day, meal) {
    return mealPlans[`${day}-${meal}`] || '';
  }
</script>

<svelte:head>
  <title>Meal Planner</title>
</svelte:head>

<div class="meal-planner">
  <h1>Weekly Meal Planner</h1>
  
  <div class="meal-grid">
    <div class="grid-header"></div>
    {#each dayLabels as dayLabel}
      <div class="grid-header">{dayLabel}</div>
    {/each}
    
    {#each meals as meal, mealIndex}
      <div class="grid-row-header">{mealLabels[mealIndex]}</div>
      {#each days as day}
        <div class="grid-cell">
          {#if loaded}
            <MealInput 
              dayOfWeek={day} 
              mealType={meal}
              initialValue={getMealValue(day, meal)}
            />
          {/if}
        </div>
      {/each}
    {/each}
  </div>
</div>