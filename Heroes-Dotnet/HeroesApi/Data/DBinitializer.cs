using HeroesApi.Models;
using System.Linq;

namespace HeroesApi.Data
{
    public static class DbInitializer
    {
        public static void Initialize(HeroContext context)
        {
            context.Database.EnsureCreated();

            if (context.HeroItems.Any())
            {
                return;
            }

            var HeroItems = new HeroItem[]
            {
            new HeroItem { Id = 1, Name = "Szymon Piotr" },
            new HeroItem { Id = 2, Name = "Andrzej" },
            new HeroItem { Id = 3, Name = "Jakub" },
            new HeroItem { Id = 4, Name = "Jan" },
            new HeroItem { Id = 5, Name = "Filip" },
            new HeroItem { Id = 6, Name = "Bartłomiej" },
            new HeroItem { Id = 7, Name = "Tomasz" },
            new HeroItem { Id = 8, Name = "Mateusz" },
            new HeroItem { Id = 9, Name = "Jakub" },
            new HeroItem { Id = 10, Name = "Tadeusz" },
            new HeroItem { Id = 11, Name = "Szymon Gorliwy" },
            new HeroItem { Id = 12, Name = "Judasz" } 
            };

            foreach (HeroItem hero in HeroItems)
            {
                context.HeroItems.Add(hero);
            }
            
            context.SaveChanges();
        }
    }
}