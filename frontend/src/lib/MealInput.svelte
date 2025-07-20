<script>
  import { searchMeals, setMealPlan } from './api.js';
  
  export let dayOfWeek;
  export let mealType;
  export let initialValue = '';
  
  let inputValue = initialValue;
  let suggestions = [];
  let showSuggestions = false;
  let searchTimeout;
  
  $: if (initialValue !== inputValue && initialValue) {
    inputValue = initialValue;
  }
  
  async function handleInput(event) {
    const value = event.target.value;
    inputValue = value;
    
    clearTimeout(searchTimeout);
    
    if (value.length > 0) {
      searchTimeout = setTimeout(async () => {
        try {
          suggestions = await searchMeals(value);
          showSuggestions = suggestions.length > 0;
        } catch (error) {
          console.error('Error searching meals:', error);
          suggestions = [];
          showSuggestions = false;
        }
      }, 300);
    } else {
      suggestions = [];
      showSuggestions = false;
    }
  }
  
  async function handleBlur() {
    setTimeout(() => {
      showSuggestions = false;
    }, 200);
    
    if (inputValue.trim()) {
      try {
        await setMealPlan(dayOfWeek, mealType, inputValue.trim());
      } catch (error) {
        console.error('Error saving meal plan:', error);
      }
    }
  }
  
  function selectSuggestion(meal) {
    inputValue = meal.name;
    showSuggestions = false;
    setMealPlan(dayOfWeek, mealType, meal.name);
  }
  
  function handleKeydown(event) {
    if (event.key === 'Enter') {
      event.target.blur();
    }
  }
</script>

<div class="autocomplete-container">
  <input
    type="text"
    class="meal-input"
    bind:value={inputValue}
    on:input={handleInput}
    on:blur={handleBlur}
    on:keydown={handleKeydown}
    placeholder="Enter meal..."
  />
  
  {#if showSuggestions && suggestions.length > 0}
    <div class="autocomplete-suggestions">
      {#each suggestions as meal}
        <div class="autocomplete-suggestion" on:click={() => selectSuggestion(meal)}>
          {meal.name}
        </div>
      {/each}
    </div>
  {/if}
</div>