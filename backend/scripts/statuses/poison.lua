count = 0

function init()
   count = math.random(1,3)
end

function tick()
   count = count-1
   p = math.random(1,5)
   creature.CurrentVitality = creature.CurrentVitality - p
   if count == 0 then
      return true
   end
   return false
end
