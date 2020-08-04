using Microsoft.EntityFrameworkCore;

namespace HeroApi.Models
{
    public class HeroContext : DbContext
    {
        public HeroContext(DbContextOptions<HeroContext> options)
            : base(options)
        {
        }

        public DbSet<HeroItem> HeroItems { get; set; }
    }
}